package utils

import (
	"reflect"
	"gopkg.in/mgo.v2/bson"
)

func T(i interface{}) string{    //函数t 有一个参数i
	switch i.(type) { //多选语句switch
	case string:
		//是字符时做的事情
	case int:
		//是整数时做的事情
	case []interface{}:
		return "array"
	}
	return ""
}

func IsInterfaceArray(i interface{}) bool {
	switch i.(type) {
	case []interface{}:
		return true
	}
	return false
}

func IsMapArray(i interface{}) bool  {
	switch i.(type) {
	case []map[string]interface{}:
	case []bson.M:
		return true
	}
	return false
}

func IsMap(i interface{}) bool {
	switch i.(type) {
	case map[string]interface{}:
		return true
	}
	return false
}

func IsBsonM(i interface{}) bool  {
	if i == nil {return false}
	return reflect.TypeOf(i).String() == "bson.M"
}

func IsString(i interface{}) bool  {
	if i == nil {return false}
	return reflect.TypeOf(i).String() == "string"
}

func IsTime(i interface{}) bool  {
	if i == nil {return false}
	return reflect.TypeOf(i).String() == "time.Time"
}

func IsArray(i interface{}) bool  {
	if i == nil {return false}
	var str = reflect.TypeOf(i).String()
	str = str[0 : 2]
	if str == "[]"{
		return true
	}
	return false
}

//func IsArray(i interface{}) bool {
//	switch i.(type) {
//	case []interface {}:
//	case []string:
//	case []int:
//	case []float64:
//	case []time.Time:
//	case []map[string]string:
//	case []map[string]interface {}:
//		return true
//	}
//	return false
//}