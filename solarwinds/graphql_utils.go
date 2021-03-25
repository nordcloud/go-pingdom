package solarwinds

import (
	"encoding/json"
	"fmt"
	"io"
)

type GraphQLRequest struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Query         string      `json:"query"`
	ResponseType  string
}

type GraphQLResponse map[string]interface{}

func NewGraphQLResponse(body io.Reader, key string) (*GraphQLResponse, error) {
	root := map[string]interface{}{}
	if err := json.NewDecoder(body).Decode(&root); err != nil {
		return nil, err
	}
	data, ok := root["data"].(map[string]interface{})
	if !ok {
		body, _ := json.Marshal(root)
		return nil, fmt.Errorf("request failed with response: %v", string(body))
	}
	graphQLResp := GraphQLResponse{}
	for k, v := range data[key].(map[string]interface{}) {
		graphQLResp[k] = v
	}
	return &graphQLResp, nil
}

func (r GraphQLResponse) isSuccess() bool {
	return r["success"].(bool)
}

func (r GraphQLResponse) message() string {
	return r["message"].(string)
}
