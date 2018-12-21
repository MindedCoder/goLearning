package main
import (
	"fmt"
	"goLearning/router"
)

func main() {
	fmt.Print("你好 go")
	rt := router.InitRouter()
	rt.Run(":8080")
}
