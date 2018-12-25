package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func IncludeObject(m bson.M, includes []string, db *mgo.Database) bson.M{
	for _, include := range includes{
		var obj = bson.M{}
		var ref = mgo.DBRef{}
		data, _:= bson.Marshal(m[include])
		bson.Unmarshal(data, &ref)
		db.FindRef(&ref).One(&obj)
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