package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
	"goLearning/utils"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(utils.Cors())
	router.OPTIONS("/1.1/functions/12kmCollectStatDatas", apis.DefaultAPI)
	router.GET("/1.1/classes/_User", apis.QueryUsers)
	router.POST("/user", apis.AddPersonAPI)
	return  router
}