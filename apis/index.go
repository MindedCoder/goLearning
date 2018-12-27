package apis

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	//"goLearning/db"
	"io/ioutil"
)

func DefaultAPI(c *gin.Context)  {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	c.JSON(http.StatusOK, gin.H{
		"status": "1234",
	})
}

func AddPersonAPI(c *gin.Context)  {
	fmt.Print("params is %+v", c.Request)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func GetUserAPI(c *gin.Context)  {
	fmt.Println("cid is", c.Request.URL.Query())
	userIds := c.Request.URL.Query()["id"]
	if userIds == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
		})
		return
	}
	userId := userIds[0]
	params := map[string]string{}
	params["className"] = "_User"
	params["id"] = userId
	//op := db.GetSessionInstance()
	//user := op.QueryObjects(map[string]string{}, nil)
	c.JSON(http.StatusOK, gin.H{
		"user": nil,
	})
}

