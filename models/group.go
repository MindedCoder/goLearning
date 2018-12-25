package models

import "gopkg.in/mgo.v2/bson"

type Group struct {
	Id  bson.ObjectId  `bson:"_id" json:"objectId"`
}
