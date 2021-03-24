package solarwinds

import (
	"encoding/json"
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
	data := root["data"].(map[string]interface{})
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
