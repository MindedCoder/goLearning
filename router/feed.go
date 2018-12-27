package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructFeedPath(router *gin.Engine)  {
	router.POST("/1.1/classes/Feed", apis.AddFeed)
	router.GET("/1.1/classes/Feed/:objectId", apis.FetchFeed)
	router.PUT("/1.1/classes/Feed/:objectId", apis.UpdateFeed)
	router.GET("/1.1/classes/Feed", apis.QueryFeeds)
	router.GET("/1.1/cloudQuery", apis.DoQuery)
	router.DELETE("/1.1/classes/Feed/:objectId", apis.DeleteFeed)
	router.GET("/1.1/scan/classes/Feed", apis.ScanFeed)
}

