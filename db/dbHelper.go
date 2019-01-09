package db

import (
	"gopkg.in/mgo.v2"
	"goLearning/utils"
	"fmt"
	"goLearning/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
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

func (op *Operater)GetDB() *mgo.Database{
	return opInstance.mgo_db
}

func (op *Operater)FetchRef(ref mgo.DBRef) bson.M {
	var obj = bson.M{}
	op.mgo_db.FindRef(&ref).One(&obj)
	return obj
}

func (op *Operater)QueryObject(queryModel models.QueryModel, params map[string]string) bson.M{
	collection := op.mgo_db.C(params["className"])
	mapInfo,_,_ := utils.Json2map(queryModel.Where, true)
	var result = bson.M{}
	m := []bson.M{
		{"$match": mapInfo},
	}
	collection.Pipe(m).One(&result)
	includes := strings.Split(queryModel.Include, ",")
	return IncludeObject(result, includes, opInstance.mgo_db)
}

func (op *Operater)QueryObjects(queryModel models.QueryModel, params map[string]string) bson.M{
	collection := op.mgo_db.C(params["className"])
	mapInfo, _, pointerKey := utils.Json2map(queryModel.Where, true)

	//查询是否含有relation字段  relation 也是通过pointer来展现
	if pointerKey != "" {
		names, _ :=op.mgo_db.CollectionNames()
		tableName := "_Join:" + mapInfo[pointerKey].(mgo.DBRef).Collection + ":" + pointerKey + ":" + params["className"]
		if IsStringInArray(names, tableName) {
			//证明是真的存在 那就得先查出该关联表中所有的params["className"]的DBRef
			relationCollection := op.mgo_db.C(tableName)
			relationRefs := []bson.M{}
			relationCollection.Find(bson.M{
				"relatedId": mapInfo[pointerKey],
			}).All(&relationRefs)
			ids := []bson.ObjectId{}
			for _, item := range relationRefs {
				owningId := item["owningId"]
				ids = append(ids, owningId.(bson.M)["$id"].(bson.ObjectId))
			}
			//再往mapInfo中放入In
			mapInfo["_id"] = bson.M{
				"$in": ids,
			}
			//再删掉之前的pointerKey
			delete(mapInfo, pointerKey)
		}
	}
	//fmt.Println("mapInfo is ", mapInfo)
	var result = []bson.M{}
	var limit = 10
	var skip = 0
	if queryModel.Limit > 0{
		limit = queryModel.Limit
	}
	if queryModel.Skip > 0{
		skip = queryModel.Skip
	}
	totalCnt := 0
	if queryModel.Count == 1{
		//说明客户端是想要计数
		totalCnt,_ = collection.Find(mapInfo).Count()

		if(queryModel.Limit == 0){
			//说明客户端只想要计数
			return bson.M{
				"count": totalCnt,
			}
		}
	}
	if queryModel.Order == ""{
		collection.Find(mapInfo).Skip(skip).Limit(limit).All(&result)
	}else {
		collection.Find(mapInfo).Skip(skip).Limit(limit).Sort(queryModel.Order).All(&result)
	}
	includes := strings.Split(queryModel.Include, ",")
	return bson.M{
		"results": IncludeObjects(result, includes, opInstance.mgo_db),
		"count": totalCnt,
	}
}

func (op *Operater) DeleteObject(params map[string]string) error{
	collection := op.mgo_db.C(params["className"])
	err := collection.RemoveId(bson.ObjectIdHex(params["objectId"]))
	return err
}

func (op *Operater) DeleteObjects(params map[string]interface{}) error {
	var objects = params["objectIds"].([]string)
	for _, id := range objects {
		var payload = map[string]string{
			"objectId": id,
			"className": params["className"].(string),
		}
		error := op.DeleteObject(payload)
		if error != nil {
			return error
			break
		}
	}
	return nil
}

func (op *Operater) AddObject(params map[string]interface{}) error  {
	collection := op.mgo_db.C(params["className"].(string))
	delete(params, "className")
	addParams := ConstructAddParams(params)
	err := collection.Insert(addParams)
	return  err
}

func (op *Operater) AddObjects(params map[string]interface{}) error  {
	var documents = params["documents"].([]map[string]string)
	for _, document := range documents {
		var payload = map[string]interface{}{
			"className": params["className"],
			"document": document,
		}
		error := op.AddObject(payload)
		if error != nil {
			return error
			break
		}
	}
	return nil
}

func (op *Operater) UpdateObject(params map[string]interface{}) error  {
	collection := op.mgo_db.C(params["className"].(string))
	var id = params["objectId"]
	delete(params, "objectId")
	delete(params, "className")
	var updateParams = ConstructUpdateParams(params)
	fmt.Println("updateparams is ", updateParams)
	err := collection.UpdateId(id, updateParams)
	return  err
}
