package solarwinds

const (
	listUserOp           = "getUsersQuery"
	listUserQuery        = "query getUsersQuery {\n  user {\n    id\n    currentOrganization {\n      id\n      members {\n        user {\n          id\n          firstName\n          lastName\n          email\n          lastLogin\n          __typename\n        }\n        role\n        products {\n          name\n          access\n          role\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"
	listUserResponseType = "user"

	getUserOp           = "getEditUserQuery"
	getUserQuery        = "query getEditUserQuery($userId: String!) {\n  user {\n    id\n    currentOrganization {\n      id\n      members(filter: {id: $userId}) {\n        id\n        user {\n          email\n          __typename\n        }\n        role\n        products {\n          name\n          role\n          access\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"
	getUserResponseType = "user"

	updateUserOp           = "updateMemberRolesMutation"
	updateUserQuery        = "mutation updateMemberRolesMutation($userId: ID!, $role: OrganizationRole!, $products: [ProductAccessInput!]) {\n  updateMemberRoles(userId: $userId, input: {role: $role, products: $products}) {\n    code\n    success\n    message\n    __typename\n  }\n}\n"
	updateUserResponseType = "updateMemberRoles"
)

type UpdateUserRequest struct {
	UserId   string          `json:"userId"`
	Role     string          `json:"role"`
	Products []ProductUpdate `json:"products"`
}

type getUserVars struct {
	UserId string `json:"userId"`
}

type UserList struct {
	OwnerUserId  string       `json:"id"`
	Organization Organization `json:"currentOrganization"`
}

type Organization struct {
	Id      string               `json:"id"`
	Members []OrganizationMember `json:"members"`
}

type OrganizationMember struct {
	User     User      `json:"user"`
	Role     string    `json:"role"`
	Products []Product `json:"products"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	LastLogin string `json:"lastLogin"`
}

type UserService struct {
	client *Client
}

func (us *UserService) List() (*UserList, error) {
	req := GraphQLRequest{
		OperationName: listUserOp,
		Query:         listUserQuery,
		ResponseType:  listUserResponseType,
	}
	resp, err := us.client.MakeGraphQLRequest(&req)
	if err != nil {
		return nil, err
	}
	userList := UserList{}
	if err := Convert(&resp, &userList); err != nil {
		return nil, err
	}
	return &userList, nil
}

func (us *UserService) Get(userId string) (*UserList, error) {
	req := GraphQLRequest{
		OperationName: getUserOp,
		Query:         getUserQuery,
		Variables: getUserVars{
			UserId: userId,
		},
		ResponseType: getUserResponseType,
	}
	resp, err := us.client.MakeGraphQLRequest(&req)
	if err != nil {
		return nil, err
	}
	userList := UserList{}
	if err := Convert(&resp, &userList); err != nil {
		return nil, err
	}
	return &userList, nil
}

func (us *UserService) Update(update UpdateUserRequest) error {
	req := GraphQLRequest{
		OperationName: updateUserOp,
		Query:         updateUserQuery,
		Variables:     update,
		ResponseType:  updateUserResponseType,
	}
	_, err := us.client.MakeGraphQLRequest(&req)
	return err
}
