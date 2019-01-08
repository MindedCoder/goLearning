package models

import (
	"gopkg.in/mgo.v2/bson"
	"goLearning/utils"
	"time"
)
/**
	DBref数组应当是断言为[]interface{}，[]map[string]interface就应该还往里面循环
	这里面有坑  待更加熟悉golang后来优化
 */
func FilterResult(m bson.M) map[string]interface{} {
	var mapInfo = map[string]interface{}{}
	for key, value := range m{
		mapInfo[key] = value
		if utils.IsBsonM(value) {
			//判断是否为DBRef
			var isDBRef = false
			for refKey, _ := range value.(bson.M) {
				if refKey == "$id" || refKey == "$ref" {
					isDBRef = true
					break
				}
			}
			if isDBRef {
				mapInfo[key] = map[string]interface{}{
					"objectId": value.(bson.M)["$id"].(bson.ObjectId).Hex(),
					"className": value.(bson.M)["$ref"].(string),
					"__type": "Pointer",
				}
			}else {
				mapInfo[key] = FilterResult(value.(bson.M))
			}
		}
		if utils.IsArray(value) {
			//是否maps数组
			if utils.IsMapArray(value) {
					var datas = []bson.M{}
					for _, item := range value.([]bson.M) {
						data := FilterResult(item)
						datas = append(datas, data)
					}
					mapInfo[key] = datas
			}else if utils.IsInterfaceArray(value) {
				//判断是否为DBRef 数组
				var isDBRef = false
				for _, item := range value.([]interface{}) {
					if utils.IsBsonM(item) {
						for refKey, _ := range item.(bson.M) {
							if refKey == "$id" || refKey == "$ref" {
								isDBRef = true
								break
							}
						}
					}
					//默认为数组里面都是同一类型
					break
				}
				if isDBRef {
					var refs = []map[string]interface{}{}
					for _, item := range value.([]interface{}) {
						ref := map[string]interface{}{
							"objectId": item.(bson.M)["$id"].(bson.ObjectId).Hex(),
							"className": item.(bson.M)["$ref"].(string),
							"__type": "Pointer",
						}
						refs = append(refs, ref)
					}
					mapInfo[key] = refs
				}
			}
		}
	}

	delete(mapInfo, "ACL")
	delete(mapInfo, "_r")
	delete(mapInfo, "_w")
	objectId := mapInfo["_id"]
	delete(mapInfo, "_id")
	if objectId != nil {
		mapInfo["objectId"] = objectId
	}
	if mapInfo["createdAt"] != nil && utils.IsTime(mapInfo["createdAt"]) {
		mapInfo["createdAt"] = mapInfo["createdAt"].(time.Time).UTC()
	}
	if mapInfo["updatedAt"] != nil && utils.IsTime(mapInfo["updatedAt"]){
		mapInfo["updatedAt"] = mapInfo["updatedAt"].(time.Time).UTC()
	}
	return mapInfo
}

func FilterResults(ms []bson.M) []bson.M {
	datas := []bson.M{}
	for _,m := range ms {
		datas = append(datas, FilterResult(m))
	}
	return datas
}