package pingdom_ext

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	defaultAuthURL = "https://my.solarwinds.cloud/v1/login"
	defaultBaseURL = "https://my.pingdom.com"
)

// Client represents a client to the Pingdom API.
type Client struct {
	JWTToken     string
	BaseURL      *url.URL
	client       *http.Client
	Integrations *IntegrationService
}

// ClientConfig represents a configuration for a pingdom client.
type ClientConfig struct {
	Username   string
	Password   string
	AuthURL    string
	BaseURL    string
	HTTPClient *http.Client
}

type authPayload struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	LoginQueryParams string `json:"loginQueryParams"`
}

// ClientConfig represents a configuration for a pingdom client.
type authResult struct {
	RedirectUrl string `json:"redirectUrl"`
}

// NewClientWithConfig returns a Pingdom client.
func NewClientWithConfig(config ClientConfig) (*Client, error) {
	var baseURL *url.URL
	var err error
	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}
	if config.AuthURL == "" {
		config.AuthURL = defaultAuthURL
	}

	baseURL, err = url.Parse(config.BaseURL)

	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL: baseURL,
	}

	if config.Username == "" {
		if envUsername, set := os.LookupEnv("PINGDOM_USERNAME"); set {
			config.Username = envUsername
		}
	}

	if config.Password == "" {
		if envPassword, set := os.LookupEnv("PINGDOM_PASSWORD"); set {
			config.Password = envPassword
		}
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	c.client = config.HTTPClient
	c.JWTToken, err = obtainToken(config)

	c.Integrations = &IntegrationService{client: c}

	return c, nil
}

func obtainToken(config ClientConfig) (string, error) {
	stateURL, err := url.Parse(config.BaseURL + "/auth/login?")
	if err != nil {
		return "", err
	}

	stateReq, err := http.NewRequest("GET", stateURL.String(), nil)
	stateResp, err := config.HTTPClient.Do(stateReq)
	if err != nil {
		return "", err
	}

	location, err := stateResp.Location()

	sessionCookie, err := getCookie(stateResp, "pingdom_login_session_id")
	if err != nil {
		return "", err
	}

	authPayload := authPayload{
		Email:            config.Username,
		Password:         config.Password,
		LoginQueryParams: location.Query().Encode(),
	}

	authBody, err := toJsonNoEscape(authPayload)
	if err != nil {
		return "", err
	}

	authReq, err := http.NewRequest("POST", config.AuthURL, bytes.NewReader(authBody))
	authReq.Header.Add("Content-Type", "application/json")

	authResp, err := config.HTTPClient.Do(authReq)
	if err != nil {
		return "", err
	}
	bodyBytes, _ := ioutil.ReadAll(authResp.Body)
	bodyString := string(bodyBytes)

	authRespJson := &authResult{}
	err1 := json.Unmarshal([]byte(bodyString), &authRespJson)

	if err1 != nil {
		return "", err1
	}

	tokenReq, err := http.NewRequest("GET", authRespJson.RedirectUrl, nil)
	tokenReq.AddCookie(sessionCookie)
	tokenResp, err := config.HTTPClient.Do(tokenReq)
	if err != nil {
		return "", err
	}

	jwtCookie, err := getCookie(tokenResp, "jwt")
	if err != nil {
		return "", err
	}

	return jwtCookie.Value, err
}

// NewRequest makes a new HTTP Request.  The method param should be an HTTP method in
// all caps such as GET, POST, PUT, DELETE.  The rsc param should correspond with
// a restful resource.  Params can be passed in as a map of strings
// Usually users of the client can use one of the convenience methods such as
// ListChecks, etc but this method is provided to allow for making other
// API calls that might not be built in.
func (pc *Client) NewRequest(method string, rsc string, params map[string]string) (*http.Request, error) {
	baseURL, err := url.Parse(pc.BaseURL.String() + rsc)
	if err != nil {
		return nil, err
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			ps.Set(k, v)
		}
		baseURL.RawQuery = ps.Encode()
	}

	req, err := http.NewRequest(method, baseURL.String(), nil)
	req.AddCookie(&http.Cookie{
		Name:  "jwt",
		Value: pc.JWTToken,
	})
	return req, err
}

// NewJSONRequest makes a new HTTP Request.  The method param should be an HTTP method in
// all caps such as GET, POST, PUT, DELETE.  The rsc param should correspond with
// a restful resource.  Params should be a json formatted string.
func (pc *Client) NewJSONRequest(method string, rsc string, params string) (*http.Request, error) {
	baseURL, err := url.Parse(pc.BaseURL.String() + rsc)
	if err != nil {
		return nil, err
	}

	reqBody := strings.NewReader(params)

	req, err := http.NewRequest(method, baseURL.String(), reqBody)
	req.AddCookie(&http.Cookie{
		Name:  "jwt",
		Value: pc.JWTToken,
	})
	req.Header.Add("Content-Type", "application/json")
	return req, err
}

// Do makes an HTTP request and will unmarshal the JSON response in to the
// passed in interface.  If the HTTP response is outside of the 2xx range the
// response will be returned along with the error.
func (pc *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return resp, err
	}

	err = decodeResponse(resp, v)
	return resp, err

}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), &v)
	return err
}

// Takes an HTTP response and determines whether it was successful.
// Returns nil if the HTTP status code is within the 2xx range.  Returns
// an error otherwise.
func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	m := &errorJSONResponse{}
	err := json.Unmarshal([]byte(bodyString), &m)
	if err != nil {
		return err
	}

	return m.Error
}

func toJsonNoEscape(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func getCookie(resp *http.Response, name string) (*http.Cookie, error) {

	for _, cookie := range resp.Cookies() {
		if cookie.Name == name {
			return cookie, nil
		}
	}

	return nil, errors.New("there is no cookie in the response")
}
