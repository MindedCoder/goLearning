package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	."goLearning/apis"
)

func main() {
	fmt.Print("你好 go")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.OPTIONS("/1.1/functions/12kmCollectStatDatas", DefaultAPI)
	//router.GET("/users", apis.GetUserAPI)
	//router.POST("/user", apis.AddPersonAPI)
	router.Run(":8080")
}
