package types

type CreateUserResponse struct {
	*APIResponse
	*User
}

type GetUserResponse struct {
	*APIResponse
	Users []*User `json:"result"`
}

type UpdateUserResponse struct {
	*APIResponse
	*User
}

type DeleteUserResponse struct {
	*APIResponse
	*User
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
