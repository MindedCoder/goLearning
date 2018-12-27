package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goLearning/models"
	"goLearning/db"
	"gopkg.in/mgo.v2/bson"
)

func QueryUsers(c *gin.Context)  {
	var queryModel models.QueryModel
	c.ShouldBindQuery(&queryModel)
	var params = map[string]string{
		"className": "_User",
	}
	oper := db.GetSessionInstance()
	result := oper.QueryObjects(queryModel, params)
	value := models.TranspileUserModels(result)
	c.JSON(http.StatusOK, gin.H{
		"result": value,
	})
}
/**
 database have no relations, so insert data for API
 */
func GetUserPermissions(c *gin.Context)  {
	params := TranspilePostParams(c)
	if params["userId"] == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1", // params error
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"editorRole": "Level0",
		"permissons": bson.M{
			"CREATE_CLASS_SITE": "CREATE_CLASS_SITE",
			"POST_SITE_NOTIFY_ARTICLE": "POST_SITE_NOTIFY_ARTICLE",
		},
	})
}
