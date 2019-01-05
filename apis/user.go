package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"goLearning/models"
	"goLearning/db"
	"gopkg.in/mgo.v2"
)

func FetchCurrentUserInfo(c *gin.Context)  {
	session := c.GetHeader("X-Lc-Session")
	oper := db.GetSessionInstance()
	mjson,_ := json.Marshal(bson.M{
		"sessionToken": session,
	})
	var queryModel = models.QueryModel{
		Where: string(mjson),
	}
	var params = map[string]string{
		"className": "_User",
	}
	user := oper.QueryObject(queryModel, params)
	result := models.FilterResult(user)
	delete(result, "password")
	c.JSON(http.StatusOK, result)
}

func FetchUserInfo(c *gin.Context)  {
	objectId := c.Param("objectId")
	if objectId == "me" {
		FetchCurrentUserInfo(c)
		return
	}
	id := c.Param("objectId")
	var queryModel models.QueryModel
	c.ShouldBindQuery(&queryModel)
	var ref = mgo.DBRef{
		Collection: "_User",
		Id: bson.ObjectIdHex(id),
	}
	oper := db.GetSessionInstance()
	object := oper.FetchRef(ref)
	if len(object) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	user := models.FilterResult(object)
	result := models.FilterResult(user)
	delete(result, "password")
	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context)  {
	postPayload := TranspilePostParams(c)
	username := postPayload["username"]
	password := postPayload["password"]
	if username == nil || password == nil {
		c.JSON(http.StatusOK, bson.M{
			"code": 1,
			"error": "Please provide username/password,mobilePhone/password or mobilePhoneNumber/smsCode",
		})
		return
	}
	oper := db.GetSessionInstance()
	mjson,_ := json.Marshal(bson.M{
		"username": username,
	})
	var queryModel = models.QueryModel{
		Where: string(mjson),
	}
	var params = map[string]string{
		"className": "_User",
	}
	user := oper.QueryObject(queryModel, params)

	if len(user) == 0 {
		c.JSON(http.StatusOK, bson.M{
			"code": 211,
			"error": "Could not find user.",
		})
		return
	}
	if user["password"] != db.Encrypt(password.(string), user["salt"].(string)) {
		c.JSON(http.StatusOK, bson.M{
			"code": 210,
			"error": "The username and password mismatch.",
		})
		return
	}
	result := models.FilterResult(user)
	delete(result, "password")
	c.JSON(http.StatusOK, result)
}

func Register(c *gin.Context) {
	//简单的注册只需要用户名以及密码就行了
	postPayload := TranspilePostParams(c)
	username := postPayload["username"]
	password := postPayload["password"]

	className := map[string]string{
		"className": "_User",
	}

	if username == nil {
		c.JSON(http.StatusOK, bson.M{
			"code": 200,
			"error": "Username is missing or empty.",
		})
		return
	}

	if password == nil {
		c.JSON(http.StatusOK, bson.M{
			"code": 201,
			"error": "Password is missing or empty.",
		})
		return
	}

	where,_ := json.Marshal(map[string]string{
		"nickname": username.(string),
	})
	//先查询nickname是否存在
	queryModel := models.QueryModel{
		Where: string(where),
	}
	oper := db.GetSessionInstance()
	user := oper.QueryObject(queryModel, className)
	if len(user) > 0 {
		//已存在nickname
		c.JSON(http.StatusOK, gin.H{
			"code": 202,
			"error": "Username has already been taken.",
		})
		return
	}
	//生成随机盐，然后生成加密密码保存
	secretPassword := db.Encrypt(password.(string), db.Salt(48, false))
	postPayload["password"] = secretPassword
	postPayload["className"] = "_User"
	postPayload["sessionToken"] = db.Salt(25, false)
	error := oper.AddObject(postPayload)

	if error == nil {
		//再查出来返回给客户端
		user = oper.QueryObject(queryModel,className)
		filterUser := models.FilterResult(user)
		c.JSON(http.StatusCreated, gin.H{
			"sessionToken": filterUser["sessionToken"],
			"createdAt": filterUser["createdAt"],
			"objectId": filterUser["objectId"],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"error": "server error",
	})
}
