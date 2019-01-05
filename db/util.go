package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"goLearning/utils"
	"crypto"
	"encoding/base64"
	"math/rand"
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
	updateParams = filterNullMap(updateParams)
	return updateParams
}

func filterNullMap(params bson.M) bson.M  {
	inc := params["$inc"].(map[string]interface{})
	set := params["$set"].(map[string]interface{})
	unset := params["$unset"].(map[string]interface{})
	addToSet := params["$addToSet"].(map[string]interface{})
	if len(inc) == 0 {
		delete(params, "$inc")
	}
	if len(set) == 0 {
		delete(params, "$set")
	}
	if len(unset) == 0 {
		delete(params, "$unset")
	}
	if len(addToSet) == 0 {
		delete(params, "$addToSet")
	}
	return params
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

func Encrypt(password string, salt string) string {
	hasher := crypto.SHA512.New()
	hasher.Write([]byte(salt))
	hasher.Write([]byte(password))
	hv := hasher.Sum(nil)
	for i := 0 ; i < 512; i++ {
		hasher.Reset()
		hasher.Write(hv)
		hv = hasher.Sum(nil)
	}
	base64String := base64.StdEncoding.EncodeToString(hv)
	return base64String
}

// Salt 生成一个盐值
func Salt(size int, needUpper bool) string {
	// 按需要生成字符串
	var result string
	var funcLength = 2
	if needUpper {
		funcLength = 3
	}
	for i := 0; i < size; i++ {
		randNumber := rand.Intn(funcLength)
		switch randNumber {
			case 0:
				result += string(Number())
				break
			case 1:
				result += string(Lower())
				break
			case 2:
				result += string(Upper())
				break
			default:
				break
		}
	}
	return result
}

// Lower 随机生成小写字母
func Lower() byte {
	lowerhouse := []int{97, 122}
	result := uint8(lowerhouse[0] + rand.Intn(26))
	return result
}
// Number 随机生成数字
func Number() byte {
	numberhouse := []int{48, 57}
	result := byte(numberhouse[0] + rand.Intn(10))
	return result
}
// Lower 随机生成大写字母
func Upper() byte {
	upperhouse := []int{65, 90}
	result := uint8(upperhouse[0] + rand.Intn(26))
	return result
}