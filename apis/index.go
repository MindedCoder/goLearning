package apis

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultAPI(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "1234",
	})
}

func AddPersonAPI(c *gin.Context)  {
	//fmt.Print("params is %+v", c.Request)
	//c.JSON(http.StatusOK, gin.H{
	//	"status": true,
	//})
}

func GetUserAPI(c *gin.Context)  {
	//fmt.Println("cid is", c.Request.URL.Query())
	//userIds := c.Request.URL.Query()["id"]
	//if userIds == nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": nil,
	//	})
	//	return
	//}
	//userId := userIds[0]
	//params := map[string]string{}
	//params["className"] = "_User"
	//params["id"] = userId
	//op := databases.GetSessionInstance()
	//user := op.GetObject(params)
	//c.JSON(http.StatusOK, gin.H{
	//	"user": user,
	//})
}
