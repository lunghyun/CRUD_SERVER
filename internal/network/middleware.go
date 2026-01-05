package network

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeoutMiddleware 모든 요청에 30초 timeout context를 설정합니다
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 기존 request의 context를 timeout context로 교체
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
