package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"strings"
)

/**
	may try loop 3 or more ,like creator.detail.school
 */
func IncludeObject(m bson.M, includes []string, db *mgo.Database) bson.M{
	//first filter includes  due to include may be "creator.detail"
	includeArray, includeMap := filterIncludes(includes)
	for _, include := range includeArray{
		var obj = bson.M{}
		var ref = mgo.DBRef{}
		data, _:= bson.Marshal(m[include])
		bson.Unmarshal(data, &ref)
		db.FindRef(&ref).One(&obj)
		for _,value := range includeMap{
			for _, subInclude := range value {
				var subObj = bson.M{}
				var subRef = mgo.DBRef{}
				subData, _:= bson.Marshal(obj[subInclude])
				bson.Unmarshal(subData, &subRef)
				db.FindRef(&subRef).One(&subObj)
				obj[subInclude] = subObj
			}
		}
		m[include] = obj
	}
	return m
}

func IncludeObjects(m []bson.M, includes []string, db *mgo.Database) []bson.M {
	var objects = []bson.M{}
	for _, obj := range m {
		objects = append(objects, IncludeObject(obj, includes, db))
	}
	return objects
}

/**
	example: ["creator", "creator.detail"]
 */
func filterIncludes(includes []string) ([]string, map[string][]string) {
	var filterArray = []string{}
	var filterMap = map[string][]string{}
	for _, include:= range includes {
		array := strings.Split(include, ".")
		filterArray = append(filterArray, array[0])
		if len(array) > 1 {
			if _, ok := filterMap[array[0]]; ok {
				valueArray := filterMap[array[0]]
				valueArray = append(valueArray, array[1])
				filterMap[array[0]] = removeDuplicatesAndEmpty(valueArray)
			}else {
				filterMap[array[0]] = []string{array[1]}
			}
		}
	}
	return removeDuplicatesAndEmpty(filterArray), filterMap
}

/**
	数组去重 去空
 */
func removeDuplicatesAndEmpty(a []string) []string{
	a_len := len(a)
	ret := []string{}
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return ret
}