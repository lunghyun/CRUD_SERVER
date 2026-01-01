package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/service"
	"github.com/lunghyun/CRUD_SERVER/types"
)

// User API 핸들러들

// 라우터
var (
	userRouterInit     sync.Once // 왜 싱글톤이냐? -> 중복 등록시 패닉됨 -> 패닉 방지
	userRouterInstance *userRouter
)

type userRouter struct {
	router      *Network
	userService *service.UserService
}

func newUserRouter(router *Network, userService *service.UserService) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}
		router.registerGET("/", userRouterInstance.get)
		router.registerDELETE("/", userRouterInstance.delete)
		router.registerCREATE("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("userRouter create")

}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("userRouter get")
	u.router.okResponse(c, &types.UserResponse{
		APIResponse: &types.APIResponse{
			Result:      1,
			Description: "success",
		},
		User: &types.User{
			Name: "user",
			Age:  24,
		},
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("userRouter update")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("userRouter delete")
}
