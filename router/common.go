package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructCommonPath(router *gin.Engine){
	router.POST("/1.1/functions/12kmFetchUserUnreadFollowCount", apis.FetchUserUnreadFollowCount)
	router.POST("/1.1/functions/12kmIsHaveNewVersion", apis.IsHaveNewVersion)
}
