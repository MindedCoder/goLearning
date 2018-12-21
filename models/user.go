package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id  bson.ObjectId  `bson:"_id" json:"id"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Avatar string   `bson:"avatar" json:"avatar"`
	Type string   `bson:"type" json:"type"`
	Detail UserDetail `bson:"detail" json:"detail"`
	Username string `bson:"username" json:"username"`
	MobilePhoneVerified bool `bson:"mobilePhoneVerified" json:"mobilePhoneVerified"`
	EmailVerified bool `bson:"emailVerified" json:"emailVerified"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type UserDetail struct {
	Id bson.ObjectId `bson:"$id" json:"id"`
}
