package main
import (
	"gopkg.in/mgo.v2"
	"goLearning/router"
)

func main() {
	mgo.SetDebug(true)
	mgo.SetStats(true)
	rt := router.InitRouter()
	rt.Run(":8080")
}
