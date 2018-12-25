package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goLearning/models"
	"goLearning/db"
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
