package controller

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goLearning/db"
	"fmt"
	"time"
)

func AddNotice()  {
	oper := db.GetSessionInstance()
	var dbref = mgo.DBRef{}
	dbref.Id = "581cbe552e958a0054decdf0"
	dbref.Collection = "_User"
	error := oper.AddObject(map[string]interface{}{
		"className":"Notice",
		"document": bson.M{
			"content":"1235643",
			"groupId":"1234567890-0987654321",
			"creator": dbref,
			"createdAt": time.Now(),
			"updatedAt": time.Now(),
		},
	})
	if error != nil {
		fmt.Print("error is ", error)
	}
}


