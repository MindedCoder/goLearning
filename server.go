package main
import (
	"goLearning/router"
	"gopkg.in/mgo.v2"
)

func main() {
	mgo.SetDebug(true)
	mgo.SetStats(true)
	rt := router.InitRouter()
	rt.Run(":5000")
}
