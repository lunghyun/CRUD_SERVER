package types

type UserResponse struct {
	*APIResponse
	*User
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
