package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"goLearning/apis"
)

func main() {
	fmt.Print("你好 go")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router.OPTIONS("/1.1/functions/12kmCollectStatDatas", apis.DefaultAPI)
	//router.GET("/users", apis.GetUserAPI)
	//router.POST("/user", apis.AddPersonAPI)
	router.Run(":8080")
}
