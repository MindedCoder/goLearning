package router

import (
	"github.com/gin-gonic/gin"
	."goLearning/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.OPTIONS("/1.1/functions/12kmCollectStatDatas", DefaultAPI)
	router.GET("/users", GetUserAPI)
	router.POST("/user", AddPersonAPI)
	return  router
}