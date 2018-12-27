package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructTestPath(router *gin.Engine)  {
	router.POST("/1.1/functions/12kmCollectStatDatas", apis.DefaultAPI)
}
