package pingdom_ext

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// IntegrationService provides an interface to Pingdom integration management.
type IntegrationService struct {
	client *Client
}

// IntegrationAPI is a Pingdom integration .
type Integration interface {
	PostParams() map[string]string
	Valid() error
}

// List returns the response holding a list of Integration.
func (cs *IntegrationService) List() ([]IntegrationGetResponse, error) {

	req, err := cs.client.NewRequest("GET", "/data/v3/integration", nil)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	m := &listIntegrationJSONResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)

	return m.Integrations, err
}

// Read returns a Integration for a given ID.
func (cs *IntegrationService) Read(id int) (*IntegrationGetResponse, error) {
	req, err := cs.client.NewRequest("GET", "/data/v3/integration/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	m := &integrationDetailsJSONResponse{}
	_, err = cs.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m.Integration, err
}

// Create a new Integration.
func (cs *IntegrationService) Create(integration Integration) (*IntegrationStatus, error) {
	if err := integration.Valid(); err != nil {
		return nil, err
	}

	req, err := cs.client.NewRequest("POST", "/data/v3/integration", integration.PostParams())
	if err != nil {
		return nil, err
	}

	m := &integrationJSONResponse{}
	_, err = cs.client.Do(req, m)
	if err != nil {
		return nil, err
	}
	return m.IntegrationStatus, err
}

// Update will update the Integration for the given ID.
func (cs *IntegrationService) Update(id int, integration Integration) (*IntegrationStatus, error) {
	if err := integration.Valid(); err != nil {
		return nil, err
	}

	req, err := cs.client.NewRequest("PUT", "/data/v3/integration/"+strconv.Itoa(id), integration.PostParams())
	if err != nil {
		return nil, err
	}

	m := &integrationJSONResponse{}
	_, err = cs.client.Do(req, m)
	if err != nil {
		return nil, err
	}
	return m.IntegrationStatus, err
}

// Delete will delete the Integration for the given ID.
func (cs *IntegrationService) Delete(id int) (*IntegrationStatus, error) {
	req, err := cs.client.NewRequest("DELETE", "/data/v3/integration/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	m := &integrationJSONResponse{}
	_, err = cs.client.Do(req, m)
	if err != nil {
		return nil, err
	}
	return m.IntegrationStatus, err
}
