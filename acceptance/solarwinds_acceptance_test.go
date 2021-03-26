package acceptance

import (
	"github.com/nordcloud/go-pingdom/solarwinds"
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
		Products: []solarwinds.ProductUpdate{
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

func TestUsers(t *testing.T) {
	if !runSolarwindsAcceptance {
		t.Skip()
	}

	userService := solarwindsClient.UserService
	currentUserEmail := os.Getenv("SOLARWINDS_USER")

	userList, err := userService.List()
	assert.NoError(t, err)
	var currentMember *solarwinds.OrganizationMember
	for _, member := range userList.Organization.Members {
		if currentUserEmail == member.User.Email {
			currentMember = &member
			break
		}
	}
	assert.True(t, currentMember != nil)

	singleUser, err := userService.Get(currentMember.User.Id)
	assert.NoError(t, err)
	assert.Equal(t, currentMember.User.Email, singleUser.Organization.Members[0].User.Email)

	containsRole := func(member *solarwinds.OrganizationMember, app string, role string) bool {
		for _, product := range member.Products {
			if product.Name == app && product.Role == role {
				return true
			}
		}
		return false
	}
	updateAddRole := solarwinds.UpdateUserRequest{
		UserId: currentMember.User.Id,
		Role:   currentMember.Role,
		Products: []solarwinds.ProductUpdate{
			{
				Name: "LOGGLY",
				Role: "MEMBER",
			},
		},
	}
	assert.True(t, containsRole(currentMember, "LOGGLY", "NO_ACCESS"))
	err = userService.Update(updateAddRole)
	assert.NoError(t, err)

	singleUser, err = userService.Get(currentMember.User.Id)
	assert.NoError(t, err)
	assert.True(t, containsRole(&singleUser.Organization.Members[0], "LOGGLY", "MEMBER"))

	updateRevokeRole := solarwinds.UpdateUserRequest{
		UserId: currentMember.User.Id,
		Role:   currentMember.Role,
		Products: []solarwinds.ProductUpdate{
			{
				Name: "LOGGLY",
				Role: "NO_ACCESS",
			},
		},
	}
	err = userService.Update(updateRevokeRole)
	assert.NoError(t, err)
	singleUser, _ = userService.Get(currentMember.User.Id)
	assert.True(t, containsRole(&singleUser.Organization.Members[0], "LOGGLY", "NO_ACCESS"))
}
