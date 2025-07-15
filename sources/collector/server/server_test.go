package server

import "github.com/gin-gonic/gin"

func createDummyGinServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	return r
}
