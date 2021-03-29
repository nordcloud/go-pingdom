package pingdom_ext

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type WebHookIntegration struct {
	Active     bool         `json:"active"`
	ProviderId int          `json:"provider_id"`
	UserData   *WebHookData `json:"user_data"`
}

type WebHookData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

/*
type LibratoIntegration struct {
	Active     bool         `json:"active"`
	ProviderId int          `json:"provider_id"`
	UserData   *LibratoData `json:"user_data"`
}

type LibratoData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ApiToken string `json:"apiToken"`
}
*/

// PostParams returns a map of parameters for an WebHook integration that can be sent along.
func (wi *WebHookIntegration) PostParams() map[string]string {
	dataJson, err := json.Marshal(wi.UserData)
	fmt.Println(err)
	m := map[string]string{
		"active":      strconv.FormatBool(wi.Active),
		"provider_id": strconv.Itoa(wi.ProviderId),
		"data_json":   string(dataJson),
	}
	return m
}

// Valid determines whether the WebHook integration contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (wi *WebHookIntegration) Valid() error {
	if wi.ProviderId != 1 && wi.ProviderId != 2 {
		return fmt.Errorf("Invalid value for `provider`.  Must contain available provider id")
	}
	if wi.UserData.Name == "" {
		return fmt.Errorf("Invalid value for `name`.  Must contain non-empty string")
	}
	if wi.UserData.Url == "" {
		return fmt.Errorf("Invalid value for `url`.  Must contain non-empty string")
	}
	return nil
}

/*

// PostParams returns a map of parameters for an Librato integration that can be sent along.
func (li *LibratoIntegration) PostParams() map[string]string {
	dataJson, err := toJsonNoEscape(li.UserData)
	fmt.Println(err)
	m := map[string]string{
		"active":      strconv.FormatBool(li.Active),
		"provider_id": strconv.Itoa(li.ProviderId),
		"data_json":   string(dataJson),
	}
	return m
}

// Valid determines whether the Librato integration contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (li *LibratoIntegration) Valid() error {
	if li.ProviderId != 1 && li.ProviderId != 2 {
		return fmt.Errorf("Invalid value for `provider`.  Must contain available provider")
	}
	if li.UserData.Name == "" {
		return fmt.Errorf("Invalid value for `name`.  Must contain non-empty string")
	}
	if li.UserData.ApiToken == "" {
		return fmt.Errorf("Invalid value for `api token`.  Must contain non-empty string")
	}
	if li.UserData.Email == "" {
		return fmt.Errorf("Invalid value for `email`.  Must contain non-empty string")
	}
	return nil
}

*/
