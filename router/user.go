package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructUserPath(router *gin.Engine)  {
	router.POST("/user", apis.AddPersonAPI)
	router.GET("/1.1/classes/_User", apis.QueryUsers)
	router.POST("/1.1/functions/12kmGetUserPermissions", apis.GetUserPermissions)
}
