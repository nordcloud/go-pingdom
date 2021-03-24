package acceptance

import (
	"github.com/chszchen-nordcloud/go-pingdom/solarwinds"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	solarwindsClient        *solarwinds.Client
	runSolarwindsAcceptance bool
)

func init() {
	if os.Getenv("SOLARWINDS_ACCEPTANCE") == "1" {
		runSolarwindsAcceptance = true
		config := solarwinds.ClientConfig{
			Username: os.Getenv("SOLARWINDS_USER"),
			Password: os.Getenv("SOLARWINDS_PASSWD"),
		}
		solarwindsClient, _ = solarwinds.NewClient(config)
		err := solarwindsClient.Init()
		if err != nil {
			panic(err)
		}
	}
}

func TestInviteUser(t *testing.T) {
	if !runSolarwindsAcceptance {
		t.Skip()
	}
	err := solarwindsClient.InvitationService.InviteUser(&solarwinds.Invitation{
		Email: solarwinds.RandString(10) + "@foo.com",
		Role:  "MEMBER",
		Products: []solarwinds.ProductSetting{
			{
				Name: "APPOPTICS",
				Role: "MEMBER",
			},
		},
	})
	assert.NoError(t, err)
}
