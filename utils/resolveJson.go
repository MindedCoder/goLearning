package utils
import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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
			bsonM["_id"] = bson.ObjectIdHex(value.(string))
		}else {
			bsonM[key] = value
		}
	}
	return bsonM, err
}

