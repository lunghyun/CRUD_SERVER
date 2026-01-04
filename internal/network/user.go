package network

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/internal/service"
	"github.com/lunghyun/CRUD_SERVER/internal/types"
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
	var req types.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil { // body로 들어오는 값을 검증, 파싱 -> req
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("바인딩 오류입니다.", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("create 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types.CreateUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1, nil),
	})
}

func (u *userRouter) get(c *gin.Context) {
	// Users:       u.userService.Get()
	user, err := u.userService.Get()
	if err != nil {
		u.router.failedResponse(c, &types.GetUserResponse{
			APIResponse: types.NewAPIResponse("조회 실패입니다", -1, err.Error()),
		})
	}
	u.router.okResponse(c, &types.GetUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1, nil),
		Users:       user,
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			APIResponse: types.NewAPIResponse("바인딩 오류입니다", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Update(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			APIResponse: types.NewAPIResponse("update 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types.UpdateUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1, nil),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			APIResponse: types.NewAPIResponse("바인딩 오류입니다", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			APIResponse: types.NewAPIResponse("delete 에러입니다", -1, err.Error()),
		})
		return
	}

	u.router.okResponse(c, &types.DeleteUserResponse{
		APIResponse: types.NewAPIResponse("성공입니다", 1, nil),
	})
}
