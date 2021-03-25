package solarwinds

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListUsers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(graphQLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		graphQLReq := GraphQLRequest{}
		json.NewDecoder(r.Body).Decode(&graphQLReq)
		assert.Equal(t, listUserOp, graphQLReq.OperationName)
		assert.Equal(t, listUserQuery, graphQLReq.Query)

		fmt.Fprintf(w, `
{
  "data": {
    "user": {
      "id": "106586091288584192",
      "currentOrganization": {
        "id": "106269109693582336",
        "members": [
          {
            "user": {
              "id": "23285292452068352",
              "firstName": "IT",
              "lastName": "Nordcloud",
              "email": "robert.kubis@nordcloud.com",
              "lastLogin": "2021-03-23T07:17:48Z",
              "__typename": "User"
            },
            "role": "ADMIN",
            "products": [
              {
                "name": "APPOPTICS",
                "access": false,
                "role": "NO_ACCESS",
                "__typename": "ProductAccess"
              },
              {
                "name": "LOGGLY",
                "access": false,
                "role": "NO_ACCESS",
                "__typename": "ProductAccess"
              },
              {
                "name": "PINGDOM",
                "access": true,
                "role": "ADMIN",
                "__typename": "ProductAccess"
              }
            ],
            "__typename": "OrganizationMember"
          },
          {
            "user": {
              "id": "74914272581727232",
              "firstName": "Nordcloud",
              "lastName": "MC-Tooling",
              "email": "mc.tooling@nordcloud.com",
              "lastLogin": "2021-03-24T23:04:56Z",
              "__typename": "User"
            },
            "role": "ADMIN",
            "products": [
              {
                "name": "APPOPTICS",
                "access": false,
                "role": "NO_ACCESS",
                "__typename": "ProductAccess"
              },
              {
                "name": "LOGGLY",
                "access": false,
                "role": "NO_ACCESS",
                "__typename": "ProductAccess"
              },
              {
                "name": "PINGDOM",
                "access": false,
                "role": "NO_ACCESS",
                "__typename": "ProductAccess"
              }
            ],
            "__typename": "OrganizationMember"
          }
        ],
        "__typename": "Organization"
      },
      "__typename": "AuthenticatedUser"
    }
  }
}
`)
	})
	userList, err := client.UserService.List()
	assert.NoError(t, err)
	members := userList.Organization.Members
	assert.Equal(t, "106586091288584192", userList.OwnerUserId)
	assert.Equal(t, len(members), 2)
}

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	userId := "106586091288584192"
	input := getUserVars{
		UserId: userId,
	}
	mux.HandleFunc(graphQLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		graphQLReq := GraphQLRequest{}
		json.NewDecoder(r.Body).Decode(&graphQLReq)
		assert.Equal(t, getUserOp, graphQLReq.OperationName)
		assert.Equal(t, getUserQuery, graphQLReq.Query)
		actualVars := getUserVars{}
		_ = Convert(&graphQLReq.Variables, &actualVars)
		assert.Equal(t, input, actualVars)
		fmt.Fprintf(w, `
{
  "data": {
    "user": {
      "id": "106586091288584192",
      "currentOrganization": {
        "id": "106269109693582336",
        "members": [
          {
            "id": "106586091288584192",
            "user": {
              "email": "chszchen@nordcloud.com",
              "__typename": "User"
            },
            "role": "ADMIN",
            "products": [
              {
                "name": "APPOPTICS",
                "role": "MEMBER",
                "access": true,
                "__typename": "ProductAccess"
              },
              {
                "name": "LOGGLY",
                "role": "NO_ACCESS",
                "access": false,
                "__typename": "ProductAccess"
              },
              {
                "name": "PINGDOM",
                "role": "ADMIN",
                "access": true,
                "__typename": "ProductAccess"
              }
            ],
            "__typename": "OrganizationMember"
          }
        ],
        "__typename": "Organization"
      },
      "__typename": "AuthenticatedUser"
    }
  }
}
`)
	})
	userList, err := client.UserService.Get("106586091288584192")
	assert.NoError(t, err)
	members := userList.Organization.Members
	assert.Equal(t, len(members), 1)
	member := members[0]
	assert.Equal(t, "chszchen@nordcloud.com", member.User.Email)
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer teardown()

	update := UpdateUserRequest{
		UserId: "106586091288584192",
		Role:   "ADMIN",
		Products: []ProductUpdate{
			{
				Name: "APPOPTICS",
				Role: "MEMBER",
			},
		},
	}
	mux.HandleFunc(graphQLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		graphQLReq := GraphQLRequest{}
		json.NewDecoder(r.Body).Decode(&graphQLReq)
		assert.Equal(t, updateUserOp, graphQLReq.OperationName)
		assert.Equal(t, updateUserQuery, graphQLReq.Query)
		actualVars := UpdateUserRequest{}
		_ = Convert(&graphQLReq.Variables, &actualVars)
		assert.Equal(t, update, actualVars)
		fmt.Fprintf(w, `
{
  "data": {
    "updateMemberRoles": {
      "code": "200",
      "success": true,
      "message": "",
      "__typename": "UpdateMemberRolesResponse"
    }
  }
}
`)
	})
	err := client.UserService.Update(update)
	assert.NoError(t, err)
}
