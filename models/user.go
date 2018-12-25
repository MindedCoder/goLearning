package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id  bson.ObjectId  `bson:"_id" json:"objectId"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Avatar string   `bson:"avatar" json:"avatar"`
	Gender int   `bson:"gender" json:"gender"`
	Type string   `bson:"type" json:"type"`
	Detail UserDetail `bson:"detail" json:"detail"`
	Email string `bson:"email" json:"email"`
	Username string `bson:"username" json:"username"`
	MobilePhoneNumber string `bson:"mobilePhoneNumber" json:"mobilePhoneNumber"`
	MobilePhoneVerified bool `bson:"mobilePhoneVerified" json:"mobilePhoneVerified"`
	EmailVerified bool `bson:"emailVerified" json:"emailVerified"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type UserDetail struct {
	Id bson.ObjectId `bson:"_id" json:"objectId"`
	LikeCnt int `bson:"likeCnt" json:"likeCnt"`
	School string `bson:"school" json:"school"`
	ClassNo string `bson:"classNo" json:"classNo"`
	UserId string `bson:"userId" json:"userId"`
	PersonalSign string `bson:"personalSign" json:"personalSign"`
	ArticleCnt int `bson:"articleCnt" json:"articleCnt"`
	BackgroundImg string `bson:"backgroundImg" json:"backgroundImg"`
	SchoolCode string `bson:"schoolCode" json:"schoolCode"`
	Campus string `bson:"campus" json:"campus"`
	Grade string `bson:"grade" json:"grade"`
	InviteCode string `bson:"inviteCode" json:"inviteCode"`
	Influence int `bson:"influence" json:"influence"`
	Rate int `bson:"rate" json:"rate"`
	HomeworkCount int `bson:"homeworkCount" json:"homeworkCount"`
	ContributionPoint int `bson:"contributionPoint" json:"contributionPoint"`
	WordCnt int `bson:"wordCnt" json:"wordCnt"`
	FeedFavCnt int `bson:"feedFavCnt" json:"feedFavCnt"`
	LikeDegree int `bson:"likeDegree" json:"likeDegree"`
	CumulativePoint int `bson:"cumulativePoint" json:"cumulativePoint"`
	AvailablePoint int `bson:"availablePoint" json:"availablePoint"`
	FollowerCnt int `bson:"followerCnt" json:"followerCnt"`
	FolloweeCnt int `bson:"followeeCnt" json:"followeeCnt"`
	FollowGroupCnt int `bson:"followGroupCnt" json:"followGroupCnt"`
	PrivacyCnt int `bson:"privacyCnt" json:"privacyCnt"`
	ReviewWordCnt int `bson:"reviewWordCnt" json:"reviewWordCnt"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	IntegralTimeFrom time.Time `bson:"integralTimeFrom" json:"integralTimeFrom"`
	IntegralTimeTo time.Time `bson:"integralTimeTo" json:"integralTimeTo"`
	LastReadTime time.Time `bson:"lastReadTime" json:"lastReadTime"`

}

func TranspileUserModel(m bson.M) User {
	data, _ := bson.Marshal(&m)
	value := User{}
	bson.Unmarshal(data, &value)
	return value
}

func TranspileUserModels(m []bson.M) []User {
	var users = []User{}
	for _, user := range m {
		transpiledUser := TranspileUserModel(user)
		users = append(users, transpiledUser)
	}
	return users
}

func TranspileUserDetailModel(m bson.M) UserDetail {
	data, _ := bson.Marshal(&m)
	value := UserDetail{}
	bson.Unmarshal(data, &value)
	return value
}

func TranspileUserdetailModels(m []bson.M) []UserDetail {
	var userdetails = []UserDetail{}
	for _, user := range m {
		transpiledUser := TranspileUserDetailModel(user)
		userdetails = append(userdetails, transpiledUser)
	}
	return userdetails
}
