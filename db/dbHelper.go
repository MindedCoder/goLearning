package databases

import (
	"gopkg.in/mgo.v2"
	"goLearning/utils"
	"goLearning/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)
var opInstance  *Operater

type Operater struct {
	mgo_session *mgo.Session
	mgo_db *mgo.Database
}

func init()  {
	infoMap := utils.Resolve()
	opInstance = new(Operater)
	opInstance.Connect(infoMap)
}

/**
	connect to databases
 */
func (oper * Operater) Connect(info map[string]string) {
	ip  := info["ip"]
	port  := info["port"]
	url := ip + ":" + port
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Source:   "admin",
		Username: info["username"],
		Password: info["password"],
	}
	mgo_session, error := mgo.DialWithInfo(dialInfo)
	oper.mgo_session = mgo_session
	oper.mgo_db = mgo_session.DB(info["databasename"])
	if error != nil {
		fmt.Println("数据库连接错误", error)
		return
	}
}

func GetSessionInstance() *Operater{
	return opInstance
}

func (op *Operater)GetObject(params map[string]string) models.User{
	collection := opInstance.mgo_db.C(params["className"])
	ps := models.User{}
	collection.Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&ps)
	return ps
}


