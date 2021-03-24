package solarwinds

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInviteUser(t *testing.T) {
	setup()
	defer teardown()

	invitation := Invitation{
		Email: RandString(8) + "@foo.com",
		Role:  "Member",
		Products: []ProductSetting{
			{
				Name: "AppOptics",
				Role: "Admin",
			},
			{
				Name: "Loggly",
				Role: "User",
			},
		},
	}
	input := inviteUserVars{
		Input: invitation,
	}
	mux.HandleFunc(graphQLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		graphQLReq := GraphQLRequest{}
		json.NewDecoder(r.Body).Decode(&graphQLReq)
		assert.Equal(t, inviteUserOp, graphQLReq.OperationName)
		assert.Equal(t, inviteUserQuery, graphQLReq.Query)
		actualVars := inviteUserVars{}
		_ = Convert(&graphQLReq.Variables, &actualVars)
		assert.Equal(t, input, actualVars)

		fmt.Fprintf(w, `
{
  "data": {
    "createOrganizationInvitation": {
      "success": true,
      "code": "200",
      "message": "",
      "invitation": {
        "email": "vB0XMNWacL@foo.com",
        "role": "MEMBER",
        "__typename": "OrganizationInvitation"
      },
      "__typename": "CreateOrganizationInvitationResponse"
    }
  }
}
`)
	})
	err := client.InvitationService.InviteUser(&invitation)
	assert.NoError(t, err)
}

func TestRevokePendingInvitation(t *testing.T) {
	setup()
	defer teardown()

	email := RandString(8) + "@foo.com"
	variables := revokeInvitationVars{
		Email: email,
	}
	mux.HandleFunc(graphQLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		graphQLReq := GraphQLRequest{}
		json.NewDecoder(r.Body).Decode(&graphQLReq)
		assert.Equal(t, revokeInvitationOp, graphQLReq.OperationName)
		assert.Equal(t, revokeInvitationQuery, graphQLReq.Query)
		actualVars := revokeInvitationVars{}
		_ = Convert(&graphQLReq.Variables, &actualVars)
		assert.Equal(t, variables, actualVars)

		fmt.Fprintf(w, `
{
  "data": {
    "deleteOrganizationInvitation": {
      "success": true,
      "code": "200",
      "message": "",
      "__typename": "MutationResponse"
    }
  }
}
`)
	})
	err := client.InvitationService.RevokePendingInvitation(email)
	assert.NoError(t, err)
}
