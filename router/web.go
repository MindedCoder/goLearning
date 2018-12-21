package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.OPTIONS("/1.1/functions/12kmCollectStatDatas", apis.DefaultAPI)
	router.GET("/users", apis.GetUserAPI)
	router.POST("/user", apis.AddPersonAPI)
	return  router
}