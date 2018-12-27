package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/utils"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(utils.Cors())
	ConstructUserPath(router)
	ConstructFeedPath(router)
	ConstructCommonPath(router)
	ConstructTestPath(router)
	return router
}