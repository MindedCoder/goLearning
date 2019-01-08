package utils
import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	ISO_TIME_FORMAT = "2006-01-02T15:04:05.999Z"
)


func Resolve() map[string]string{
	value := make(map[string]string)
	b ,err := ioutil.ReadFile("DBConfig.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(b,&value)
	return value
}

func Json2map(str string, filterId bool) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	if !filterId {
		return result, nil
	}
	bsonM := bson.M{}
	for key, value := range result {
		if IsMap(value) {
			if key == "objectId" {
				// $in
				for objKey, objValue := range value.(map[string]interface{}) {
					if objKey == "$in" {
						var objectIds = []bson.ObjectId{}
						for _, id := range objValue.([]interface{}) {
							objectIds = append(objectIds, bson.ObjectIdHex(id.(string)))
						}
						bsonM["_id"] = bson.M{
							"$in": objectIds,
						}
					}
				}
			}else {
				if value.(map[string]interface{})["__type"] == "Pointer" {
					bsonM[key] = mgo.DBRef{
						Id: bson.ObjectIdHex(value.(map[string]interface{})["objectId"].(string)),
						Collection: value.(map[string]interface{})["className"].(string),
					}
				}else {
					//需要注意的是时间格式 客户端传过来的是createdAt:map[$lt:map[__type:Date iso:2019-01-08T08:00:24.240Z]]
					timeLessMap := value.(map[string]interface{})["$lt"]
					timeGreaterMap := value.(map[string]interface{})["$gt"]
					if timeLessMap != nil {
						t,_ := time.Parse(ISO_TIME_FORMAT, timeLessMap.(map[string]interface{})["iso"].(string))
						value.(map[string]interface{})["$lt"] = time.Time.Local(t)
						bsonM[key] = value
					}
					if timeGreaterMap != nil {
						t,_ := time.Parse(ISO_TIME_FORMAT, timeGreaterMap.(map[string]interface{})["iso"].(string))
						value.(map[string]interface{})["$gt"] = time.Time.Local(t)
						bsonM[key] = value
					}
					bsonM[key] = value
				}
			}
		}else {
			if key == "objectId" {
				bsonM["_id"] = bson.ObjectIdHex(value.(string))
			}else {
				bsonM[key] = value
			}
		}
	}
	return bsonM, err
}

