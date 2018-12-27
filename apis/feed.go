package apis

import (
	"github.com/gin-gonic/gin"
	"goLearning/models"
	"fmt"
	"net/http"
	"goLearning/db"
)

func AddFeed(c *gin.Context)  {

}

func DeleteFeed(c *gin.Context)  {

}

func UpdateFeed(c *gin.Context)  {

}

func FetchFeed(c *gin.Context)  {

}

func QueryFeeds(c *gin.Context)  {
	var queryModel models.QueryModel
	c.ShouldBindQuery(&queryModel)
	fmt.Print("feed Search model is ", queryModel)
	var params = map[string]string{
		"className": "Feed",
	}
	oper := db.GetSessionInstance()
	result := oper.QueryObjects(queryModel, params)
	value := models.TranspileFeedModels(result)
	c.JSON(http.StatusOK, gin.H{
		"results": value,
	})
}

func DoQuery(c *gin.Context)  {

}
func ScanFeed(c *gin.Context)  {

}