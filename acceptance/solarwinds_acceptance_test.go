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

func TestInvitations(t *testing.T) {
	if !runSolarwindsAcceptance {
		t.Skip()
	}
	email := solarwinds.RandString(10) + "@foo.com"
	invitationService := solarwindsClient.InvitationService
	err := invitationService.InviteUser(&solarwinds.Invitation{
		Email: email,
		Role:  "MEMBER",
		Products: []solarwinds.ProductSetting{
			{
				Name: "APPOPTICS",
				Role: "MEMBER",
			},
		},
	})
	assert.NoError(t, err)

	err = invitationService.ResendInvitation(email)
	assert.NoError(t, err)

	err = invitationService.RevokePendingInvitation(email)
	assert.NoError(t, err)

	err = invitationService.ResendInvitation(email)
	assert.Error(t, err)
}
