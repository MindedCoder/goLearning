package router

import (
	"github.com/gin-gonic/gin"
	"goLearning/apis"
)

func ConstructCommonPath(router *gin.Engine){
	router.POST("/1.1/classes/:object", apis.AddObject)
	router.GET("/1.1/classes/:object/:objectId", apis.FetchObject)
	router.PUT("/1.1/classes/:object/:objectId", apis.UpdateObject)
	router.GET("/1.1/classes/:object", apis.QueryObjects)
	router.OPTIONS("/1.1/classes/:object", apis.QueryObjects)
	router.DELETE("/1.1/classes/:object/:objectId", apis.DeleteObject)
	router.POST("/1.1/batch", apis.Batch)
}
