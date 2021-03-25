package solarwinds

const (
	inviteUserOp           = "createOrganizationAdminMutation"
	inviteUserQuery        = "mutation createOrganizationAdminMutation($input: CreateOrganizationInvitationInput!) {\n  createOrganizationInvitation(input: $input) {\n    success\n    code\n    message\n    invitation {\n      email\n      role\n      __typename\n    }\n    __typename\n  }\n}\n"
	inviteUserResponseType = "createOrganizationInvitation"

	revokeInvitationOp           = "deleteOrganizationInvitationMutation"
	revokeInvitationQuery        = "mutation deleteOrganizationInvitationMutation($email: ID!) {\n  deleteOrganizationInvitation(email: $email) {\n    success\n    code\n    message\n    __typename\n  }\n}\n"
	revokeInvitationResponseType = "deleteOrganizationInvitation"

	resendInvitationOp           = "resendOrganizationInvitationMutation"
	resendInvitationQuery        = "mutation resendOrganizationInvitationMutation($email: ID!) {\n  resendOrganizationInvitation(email: $email) {\n    success\n    code\n    message\n    __typename\n  }\n}\n"
	resendInvitationResponseType = "resendOrganizationInvitation"
)

type Invitation struct {
	Email    string          `json:"email"`
	Role     string          `json:"role"`
	Products []ProductUpdate `json:"products"`
}

type inviteUserVars struct {
	Input Invitation `json:"input"`
}

type revokeInvitationVars struct {
	Email string `json:"email"`
}

type resendInvitationVars struct {
	Email string `json:"email"`
}

type InvitationService struct {
	client *Client
}

func (is *InvitationService) InviteUser(user *Invitation) error {
	req := GraphQLRequest{
		OperationName: inviteUserOp,
		Query:         inviteUserQuery,
		Variables: inviteUserVars{
			Input: *user,
		},
		ResponseType: inviteUserResponseType,
	}
	_, err := is.client.MakeGraphQLRequest(&req)
	return err
}

func (is *InvitationService) RevokePendingInvitation(email string) error {
	req := GraphQLRequest{
		OperationName: revokeInvitationOp,
		Query:         revokeInvitationQuery,
		Variables: revokeInvitationVars{
			Email: email,
		},
		ResponseType: revokeInvitationResponseType,
	}
	_, err := is.client.MakeGraphQLRequest(&req)
	return err
}

func (is *InvitationService) ResendInvitation(email string) error {
	req := GraphQLRequest{
		OperationName: resendInvitationOp,
		Query:         resendInvitationQuery,
		Variables: resendInvitationVars{
			Email: email,
		},
		ResponseType: resendInvitationResponseType,
	}
	_, err := is.client.MakeGraphQLRequest(&req)
	return err
}
