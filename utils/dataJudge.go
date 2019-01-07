package utils

import "reflect"

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

func IsArray(i interface{}) bool {
	switch i.(type) {
	case []interface{}:
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
	return reflect.TypeOf(i).String() == "bson.M"
}

func IsString(i interface{}) bool  {
	return reflect.TypeOf(i).String() == "string"
}

func IsTime(i interface{}) bool  {
	return reflect.TypeOf(i).String() == "time.Time"
}
