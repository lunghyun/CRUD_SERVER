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

	err := u.userService.Create(&types.User{
		Name: "user",
		Age:  24,
	})

	if err != nil {
		fmt.Println(err)
	}

	u.router.okResponse(c, &types.CreateUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("userRouter get")

	u.router.okResponse(c, &types.GetUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("userRouter update")

	err := u.userService.Update(nil, nil)

	if err != nil {
		fmt.Println(err)
	}

	u.router.okResponse(c, &types.UpdateUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("userRouter delete")

	err := u.userService.Delete(nil)

	if err != nil {
		fmt.Println(err)
	}

	u.router.okResponse(c, &types.DeleteUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1),
	})
}
