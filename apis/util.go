package apis

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

func TranspilePostParams(c *gin.Context) bson.M {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var params = bson.M{}
	bson.UnmarshalJSON(data, &params)
	return params
}