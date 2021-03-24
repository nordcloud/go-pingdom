package solarwinds

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

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
