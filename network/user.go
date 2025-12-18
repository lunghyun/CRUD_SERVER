package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/types"
)

// 라우터
var (
	userRouterInit     sync.Once // 1번만 호출 되어야함
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}
		router.registerGET("/", userRouterInstance.get)
		router.registerDELETE("/", userRouterInstance.delete)
		router.registerCREATE("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) { // gin을 사용할때는 api라는 걸 명시하기 위해 context를 사용해야한다?
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
