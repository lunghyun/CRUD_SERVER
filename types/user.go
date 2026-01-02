package types

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserResponse struct {
	*APIResponse
	Users []*User `json:"result"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"` // binding 규칙에 어긋나면 에러 노출
	Age  int    `json:"age" binding:"required"`
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*APIResponse
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func (c *UpdateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type UpdateUserResponse struct {
	*APIResponse
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type DeleteUserResponse struct {
	*APIResponse
}
