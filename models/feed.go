package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feed struct {
	Id 				      bson.ObjectId  `bson:"_id" json:"objectId"`
	LikeCount       int 					 `bson:"likeCount" json:"likeCount"`
	RewardCount     int 					 `bson:"rewardCount" json:"rewardCount"`
	ReviewCount     int 					 `bson:"reviewCount" json:"reviewCount"`
	CommentCount    int 					 `bson:"commentCount" json:"commentCount"`
	TotalSocialCount     int 			 `bson:"totalSocialCount" json:"totalSocialCount"`
	TotalCommentCount    int 			 `bson:"totalCommentCount" json:"totalCommentCount"`
	DirectCommentCount   int 			 `bson:"directCommentCount" json:"directCommentCount"`
	LikeDegree      int 			     `bson:"likeDegree" json:"likeDegree"`
	ReadCount       int 					 `bson:"readCount" json:"readCount"`
	WordCount       int 					 `bson:"wordCount" json:"wordCount"`
	FavCount        int 					 `bson:"favCount" json:"favCount"`
	RepostCount     int 					 `bson:"repostCount" json:"repostCount"`
	Active          int 					 `bson:"active" json:"active"`
	TotalScore      int 					 `bson:"totalScore" json:"totalScore"`
	Score           float64 			 `bson:"score" json:"score"`
	Tags 			      []string 		   `bson:"tags" json:"tags"`
	InnerTags 			[]string 		   `bson:"innerTags" json:"innerTags"`
	Articles 			  []Article 		 `bson:"articles" json:"articles"`
	Article 			  Article 		   `bson:"article" json:"article"`
	Creator 	      User					 `bson:"creator" json:"creator"`
	LikeUsers 	    []User				 `bson:"likeUsers" json:"likeUsers"`
	PostedGroups 	  []Group				 `bson:"postedGroups" json:"postedGroups"`
	ContributeGroups     []Group	 `bson:"contributeGroups" json:"contributeGroups"`
	Cover 	        string				 `bson:"cover" json:"cover"`
	Abstract 	      string				 `bson:"abstract" json:"abstract"`
	Title 	        string				 `bson:"title" json:"title"`
	FeedType 	      string				 `bson:"feedType" json:"feedType"`
	Category 	      string				 `bson:"category" json:"category"`
	CtrbTags 	      []string			 `bson:"ctrbTags" json:"ctrbTags"`
	Privacy 	      bool				   `bson:"privacy" json:"privacy"`
	PubTime         time.Time      `bson:"pubTime" json:"pubTime"`
	ReplyTime       time.Time      `bson:"replyTime" json:"replyTime"`
	CreatedAt       time.Time      `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time      `bson:"updatedAt" json:"updatedAt"`
}

type Article struct {
	Id 				      bson.ObjectId  `bson:"_id" json:"objectId"`
} 