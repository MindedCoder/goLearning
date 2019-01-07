package utils
import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
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
		if key == "objectId" {
			if !IsMap(value){
				bsonM["_id"] = bson.ObjectIdHex(value.(string))
			}else {
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
			}
		}else if IsMap(value) {
			if value.(map[string]interface{})["__type"] == "Pointer" {
				bsonM[key] = mgo.DBRef{
					Id: bson.ObjectIdHex(value.(map[string]interface{})["objectId"].(string)),
					Collection: value.(map[string]interface{})["className"].(string),
				}
			}
 		}else {
			bsonM[key] = value
		}
	}
	return bsonM, err
}

