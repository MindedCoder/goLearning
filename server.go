package main
import (
	."goLearning/router"
	"gopkg.in/mgo.v2"
)

func main() {
	mgo.SetDebug(true)
	mgo.SetStats(true)
	router := InitRouter()
	router.Run(":8080")
}
