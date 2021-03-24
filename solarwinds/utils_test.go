package solarwinds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJsonNoEscape(t *testing.T) {
	obj := map[string]interface{}{
		"url":     "https://foo.bar.com?name=a&value=b",
		"content": "<html>response body</html>",
	}
	b, err := toJsonNoEscape(obj)
	assert.NoError(t, err)
	assert.Equal(t, `{"content":"<html>response body</html>","url":"https://foo.bar.com?name=a&value=b"}`+"\n", string(b))
}

func TestRandString(t *testing.T) {
	size := 10
	for i := 5; i > 0; i -= 1 {
		a, b := RandString(size), RandString(size)
		assert.Equal(t, size, len(a))
		assert.Equal(t, size, len(b))
		assert.True(t, a != b)
	}
}

func TestConvert(t *testing.T) {
	type Bar struct {
		Name string `json:"name"`
	}

	type Foo struct {
		Bar Bar `json:"bar"`
	}

	dynObj := map[string]interface{}{
		"bar": map[string]interface{}{
			"name": "this is bar in foo",
		},
	}
	actual := Foo{}
	_ = Convert(&dynObj, &actual)
	expected := Foo{
		Bar: Bar{
			Name: "this is bar in foo",
		},
	}
	assert.Equal(t, expected, actual)
}
