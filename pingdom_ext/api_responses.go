package pingdom_ext

import (
	"github.com/nordcloud/go-pingdom/pingdom"
)

type errorJSONResponse struct {
	Error *pingdom.PingdomError `json:"error"`
}

type listIntegrationJSONResponse struct {
	Integrations []IntegrationGetResponse `json:"integration"`
}

type integrationDetailsJSONResponse struct {
	Integration *IntegrationGetResponse `json:"integration"`
}

type integrationJSONResponse struct {
	IntegrationStatus *IntegrationStatus `json:"integration"`
}

// Integration Get Response
type IntegrationGetResponse struct {
	NumberOfConnectedChecks int               `json:"number_of_connected_checks"`
	Id                      int               `json:"id"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	ProviderId              int               `json:"provider_id"`
	ActivatedAt             int               `json:"activated_at"`
	CreatedAt               int               `json:"created_at"`
	UserData                map[string]string `json:"user_data"`
}

type IntegrationStatus struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}
