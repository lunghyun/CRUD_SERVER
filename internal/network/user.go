package network

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/internal/service"
	types2 "github.com/lunghyun/CRUD_SERVER/internal/types"
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
	var req types2.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil { // body로 들어오는 값을 검증, 파싱 -> req
		u.router.failedResponse(c, &types2.CreateUserResponse{
			APIResponse: types2.NewAPIResponse("바인딩 오류입니다.", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types2.CreateUserResponse{
			APIResponse: types2.NewAPIResponse("create 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types2.CreateUserResponse{
		APIResponse: types2.NewAPIResponse("성공입니다", 1, nil),
	})
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types2.GetUserResponse{
		APIResponse: types2.NewAPIResponse("성공입니다", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types2.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types2.UpdateUserResponse{
			APIResponse: types2.NewAPIResponse("바인딩 오류입니다", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Update(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types2.UpdateUserResponse{
			APIResponse: types2.NewAPIResponse("update 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types2.UpdateUserResponse{
		APIResponse: types2.NewAPIResponse("성공입니다", 1, nil),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	var req types2.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types2.DeleteUserResponse{
			APIResponse: types2.NewAPIResponse("바인딩 오류입니다", -1, err.Error()),
		})
	}

	if err := u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types2.DeleteUserResponse{
			APIResponse: types2.NewAPIResponse("delete 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types2.DeleteUserResponse{
		APIResponse: types2.NewAPIResponse("성공입니다", 1, nil),
	})
}
