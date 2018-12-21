package main
import (
	"goLearning/router"
	"gopkg.in/mgo.v2"
	"fmt"
)

func main() {
	mgo.SetDebug(true)
	mgo.SetStats(true)
	rt := router.InitRouter()
	fmt.Println("Listening and serving HTTP on :8080")
	rt.Run(":8080")
}
