package acceptance

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/nordcloud/go-pingdom/pingdom_ext"
	"github.com/stretchr/testify/assert"
)

var client_ext *pingdom_ext.Client

var runExtAcceptance bool

func init() {
	if os.Getenv("PINGDOM_ACCEPTANCE") == "1" {
		runExtAcceptance = true

		config := pingdom_ext.ClientConfig{
			HTTPClient: &http.Client{
				Timeout: time.Second * 10,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			},
		}
		client_ext, _ = pingdom_ext.NewClientWithConfig(config)

	}
}

func TestIntegrations(t *testing.T) {
	if !runExtAcceptance {
		t.Skip()
	}

	integration := pingdom_ext.WebHookIntegration{
		Active:     false,
		ProviderId: 2,
		UserData: &pingdom_ext.WebHookData{
			Name: "wlwu-tets-1",
			Url:  "http://www.example.com",
		},
	}

	createMsg, err := client_ext.Integrations.Create(&integration)
	assert.NoError(t, err)
	assert.NotNil(t, createMsg)
	assert.NotEmpty(t, createMsg)

	integrationID := createMsg.ID

	listMsg, err := client_ext.Integrations.List()
	assert.NoError(t, err)
	assert.NotNil(t, listMsg)
	assert.NotEmpty(t, listMsg)

	getMsg, err := client_ext.Integrations.Read(integrationID)
	assert.NoError(t, err)
	assert.NotNil(t, getMsg)
	assert.NotEmpty(t, getMsg)
	assert.NotEmpty(t, getMsg.CreatedAt)
	assert.NotEmpty(t, getMsg.Name)

	integration.Active = true
	integration.UserData.Name = "wlwu-tets-update"
	integration.UserData.Url = "http://www.example1.com"

	updateMsg, err := client_ext.Integrations.Update(integrationID, &integration)
	assert.NoError(t, err)
	assert.NotNil(t, updateMsg)

	delMsg, err := client_ext.Integrations.Delete(integrationID)
	assert.NoError(t, err)
	assert.NotNil(t, delMsg)

}
