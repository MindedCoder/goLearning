package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"goLearning/utils"
)

func ConstructUpdateParams(params bson.M) bson.M {
	var updateParams = bson.M{
		"$set": map[string]interface{}{},
		"$inc": map[string]interface{}{},
		"$unset": map[string]interface{}{},
		"$addToSet": map[string]interface{}{},
	}
	for key, value := range params {
		if !utils.IsMap(value){
			//是否考虑更新时候全覆盖的情况？
			//这里还要判断数组里面是否为DBref
			if utils.IsArray(value){
				var isMapArray = false
				for _, item := range value.([]interface{}) {
					if utils.IsMap(item) {
						isMapArray = true
						break
					}
				}
				if(isMapArray){
					//这个时候应当是DBref
					var refs = []mgo.DBRef{}
					for _, object := range value.([]interface{}) {
						//这里需要注意object的__type 暂时忽略
						var ref = mgo.DBRef{
							Collection: object.(map[string]interface{})["className"].(string),
							Id: bson.ObjectIdHex(object.(map[string]interface{})["objectId"].(string)),
						}
						refs = append(refs, ref)
					}
					updateParams["$set"].(map[string]interface{})[key] = refs
				}else {
					updateParams["$set"].(map[string]interface{})[key] = value
				}
			}else {
				updateParams["$set"].(map[string]interface{})[key] = value
			}

		}else {
			var _type = value.(map[string]interface{})["__type"]
			if _type == "Pointer" {
				updateParams["$set"].(map[string]interface{})[key] = mgo.DBRef{
					Collection: value.(map[string]interface{})["className"].(string),
					Id: bson.ObjectIdHex(value.(map[string]interface{})["objectId"].(string)),
				}
			} else {
				var op = value.(map[string]interface{})["__op"]
				if op == "Increment"{
					updateParams["$inc"].(map[string]interface{})[key] = int(value.(map[string]interface{})["amount"].(float64))
				}
				if op == "Delete"{
					updateParams["$unset"].(map[string]interface{})[key] = ""
				}
				if op == "AddUnique"{
					var objects = value.(map[string]interface{})["objects"].([]interface{})
					var refs = []mgo.DBRef{}
					for _, object := range objects {
						//这里需要注意object的__type 暂时忽略
						var ref = mgo.DBRef{
							Collection: object.(map[string]interface{})["className"].(string),
							Id: bson.ObjectIdHex(object.(map[string]interface{})["objectId"].(string)),
						}
						refs = append(refs, ref)
					}
					updateParams["$addToSet"].(map[string]interface{})[key] = bson.M{
						"$each": refs,
					}
				}
			}
		}
	}
	return updateParams
}



func ConstructAddParams(params bson.M) bson.M {
	var addParams = bson.M{}
	for key, value := range params {
		if !utils.IsMap(value){
			//这里还要判断数组里面是否为DBref
			if utils.IsArray(value){
				var isMapArray = false
				for _, item := range value.([]interface{}) {
					if utils.IsMap(item) {
						isMapArray = true
						break
					}
				}
				if(isMapArray){
				 //这个时候应当是DBref
					var refs = []mgo.DBRef{}
					for _, object := range value.([]interface{}) {
						//这里需要注意object的__type 暂时忽略
						var ref = mgo.DBRef{
							Collection: object.(map[string]interface{})["className"].(string),
							Id: bson.ObjectIdHex(object.(map[string]interface{})["objectId"].(string)),
						}
						refs = append(refs, ref)
					}
					addParams[key] = refs
				}else {
					addParams[key] = value
				}
			}else {
				addParams[key] = value
			}
		}else {
			var _type = value.(map[string]interface{})["__type"]
			if _type == "Pointer" {
				addParams[key] = mgo.DBRef{
					Collection: value.(map[string]interface{})["className"].(string),
					Id: bson.ObjectIdHex(value.(map[string]interface{})["objectId"].(string)),
				}
			} else {
				var op = value.(map[string]interface{})["__op"]
				if op == "AddUnique" {
					var objects = value.(map[string]interface{})["objects"].([]interface{})
					var refs = []mgo.DBRef{}
					for _, object := range objects {
						//这里需要注意object的__type 暂时忽略
						var ref = mgo.DBRef{
							Collection: object.(map[string]interface{})["className"].(string),
							Id: bson.ObjectIdHex(object.(map[string]interface{})["objectId"].(string)),
						}
						refs = append(refs, ref)
					}
					addParams[key] = refs
				}
			}
		}
	}
	var timeNow = time.Now()
	addParams["createdAt"] = timeNow
	addParams["updatedAt"] = timeNow
	return addParams
}