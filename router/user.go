package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructUserPath(router *gin.Engine)  {
	router.GET("/1.1/users/:objectId", apis.FetchUserInfo)
	router.POST("/1.1/login", apis.Login)
	router.POST("/1.1/users", apis.Register)
}
