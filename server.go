package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("你好 go")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	fmt.Println(router)
	//router.OPTIONS("/1.1/functions/12kmCollectStatDatas", apis.DefaultAPI)
	//router.GET("/users", apis.GetUserAPI)
	//router.POST("/user", apis.AddPersonAPI)
	//router.Run(":8080")
}
