package solarwinds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func toJsonNoEscape(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func retrieveCookie(resp *http.Response, name string) (string, error) {
	if cookies := resp.Header[headerNameSetCookie]; cookies != nil {
		for _, cookie := range cookies {
			if strings.Contains(cookie, name+"=") {
				if parts := strings.Split(cookie, ";"); parts != nil {
					for _, part := range parts {
						if strings.HasPrefix(part, name) {
							return strings.Split(part, "=")[1], nil
						}
					}
				}
			}
		}
		return "", fmt.Errorf("cookie '%v' does not exist in the response", name)
	} else {
		return "", errors.New("there is no cookie in the response")
	}
}
