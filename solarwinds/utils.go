package solarwinds

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Note there will be a new line character at the end of the output.
func toJsonNoEscape(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// For testing purpose only.
func Convert(from interface{}, to interface{}) error {
	b, err := toJsonNoEscape(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, to)
}
