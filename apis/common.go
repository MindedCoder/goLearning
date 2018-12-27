package apis

import (
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	//"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func FetchUserUnreadFollowCount(c *gin.Context)  {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("ctx.Request.body: %v", string(data))
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
	})
}

func IsHaveNewVersion(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"result": bson.M{
			"canUpdate": false,
			"updateDesc": "12KM正在系统升级中，预计8月2日9时升级完成。敬请期待！",
			"forceUpdate": false,
			"url": "https://www.12km.com/download",
		},
	})
}
