package pingdom_ext

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClientWithConfig(t *testing.T) {
	type args struct {
		config ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientWithConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClientWithConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_obtainToken(t *testing.T) {
	type args struct {
		config ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := obtainToken(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("obtainToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("obtainToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		JWTToken     string
		BaseURL      *url.URL
		client       *http.Client
		Integrations *IntegrationService
	}
	type args struct {
		method string
		rsc    string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &Client{
				JWTToken:     tt.fields.JWTToken,
				BaseURL:      tt.fields.BaseURL,
				client:       tt.fields.client,
				Integrations: tt.fields.Integrations,
			}
			got, err := pc.NewRequest(tt.args.method, tt.args.rsc, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewJSONRequest(t *testing.T) {
	type fields struct {
		JWTToken     string
		BaseURL      *url.URL
		client       *http.Client
		Integrations *IntegrationService
	}
	type args struct {
		method string
		rsc    string
		params string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &Client{
				JWTToken:     tt.fields.JWTToken,
				BaseURL:      tt.fields.BaseURL,
				client:       tt.fields.client,
				Integrations: tt.fields.Integrations,
			}
			got, err := pc.NewJSONRequest(tt.args.method, tt.args.rsc, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewJSONRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewJSONRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		JWTToken     string
		BaseURL      *url.URL
		client       *http.Client
		Integrations *IntegrationService
	}
	type args struct {
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &Client{
				JWTToken:     tt.fields.JWTToken,
				BaseURL:      tt.fields.BaseURL,
				client:       tt.fields.client,
				Integrations: tt.fields.Integrations,
			}
			got, err := pc.Do(tt.args.req, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeResponse(t *testing.T) {
	type args struct {
		r *http.Response
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeResponse(tt.args.r, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateResponse(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_toJsonNoEscape(t *testing.T) {
	type args struct {
		t interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toJsonNoEscape(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("toJsonNoEscape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toJsonNoEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCookie(t *testing.T) {
	type args struct {
		resp *http.Response
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Cookie
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCookie(tt.args.resp, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
