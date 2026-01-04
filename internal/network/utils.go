package network

import (
	"github.com/gin-gonic/gin"
)

// register 유틸 함수들
func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerCREATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

// response 형식 맞추는 유틸 함수
func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) createResponse(c *gin.Context, result interface{}) {
	c.JSON(201, result)
}

func (n *Network) failedResponse(c *gin.Context, status int, result interface{}) {
	c.JSON(status, result)
}
