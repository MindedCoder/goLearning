package models

import (
	"gopkg.in/mgo.v2/bson"
	"goLearning/utils"
	"time"
)

func FilterResult(m bson.M) map[string]interface{} {
	delete(m, "ACL")
	delete(m, "_r")
	delete(m, "_w")
	objectId := m["_id"]
	delete(m, "_id")
	if objectId != nil {
		m["objectId"] = objectId
	}
	if m["createdAt"] != nil {
		m["createdAt"] = m["createdAt"].(time.Time).UTC()
	}

	if m["updatedAt"] != nil {
		m["updatedAt"] = m["updatedAt"].(time.Time).UTC()
	}

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
			//判断是否为DBRef 数组
			var isDBRef = false
			var isBsonM = false
			for _, item := range value.([]interface{}) {
				if utils.IsBsonM(item) {
					isBsonM = true
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
			}else {
				if isBsonM {
					var datas = []map[string]interface{}{}
					for _, item := range value.([]interface{}) {
						data := FilterResult(item.(bson.M))
						datas = append(datas, data)
					}
					mapInfo[key] = datas
				}
			}
		}
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