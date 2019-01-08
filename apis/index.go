package apis

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"goLearning/db"
	"goLearning/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"fmt"
	"time"
)

func AddObject(c *gin.Context)  {
	className := c.Param("object")
	params := TranspilePostParams(c)
	objectId := bson.NewObjectId()
	params["className"] = className
	params["_id"] = objectId
	oper := db.GetSessionInstance()
	error := oper.AddObject(params)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"error": "server error",
		})
		return
	}
	var ref = mgo.DBRef{
		Collection: className,
		Id: objectId,
	}
	object := oper.FetchRef(ref)
	object = models.FilterResult(object)
	c.Header("Location", "/1.1/classes/" + className + "/" + objectId.String())//暂时未知URL咋写,前面得拼一个域名
	fetchWhenSave := c.Query("fetchWhenSave")
	if fetchWhenSave == "true"{
		c.JSON(http.StatusCreated, object)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"objectId": object["objectId"],
		"createdAt": object["createdAt"],
	})
}

func DeleteObject(c *gin.Context)  {
	id := c.Param("objectId")
	className := c.Param("object")
	var params = map[string]string{
		"objectId": id,
		"className": className,
	}
	oper := db.GetSessionInstance()
	error := oper.DeleteObject(params)
	if error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func UpdateObject(c *gin.Context)  {
	id := c.Param("objectId")
	className := c.Param("object")
	params := TranspilePostParams(c)
	params["objectId"] = bson.ObjectIdHex(id)
	params["className"] = className
	params["updatedAt"] = time.Now()
	oper := db.GetSessionInstance()
	error := oper.UpdateObject(params)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"error": "server error",
		})
		return
	}

	fetchWhenSave := c.Query("fetchWhenSave")
	if fetchWhenSave != "true"{
		c.JSON(http.StatusOK, gin.H{
			"updatedAt": params["updatedAt"],
		})
		return
	}
	//若有fetchWhenSave 也仅仅返回被更新的字段
	c.JSON(http.StatusOK, params)
}

func FetchObject(c *gin.Context)  {
	id := c.Param("objectId")
	className := c.Param("object")
	var queryModel models.QueryModel
	c.ShouldBindQuery(&queryModel)
	var ref = mgo.DBRef{
		Collection: className,
		Id: bson.ObjectIdHex(id),
	}
	oper := db.GetSessionInstance()
	object := oper.FetchRef(ref)
	if len(object) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	result := db.IncludeObject(object, strings.Split(queryModel.Include, ","), oper.GetDB())
	value := models.FilterResult(result)
	c.JSON(http.StatusOK, value)
}

func QueryObjects(c *gin.Context)  {
	className := c.Param("object")
	var queryModel models.QueryModel
	c.ShouldBindQuery(&queryModel)
	var params = map[string]string{
		"className": className,
	}
	oper := db.GetSessionInstance()
	result := oper.QueryObjects(queryModel, params)

	if queryModel.Count == 1 {
		c.JSON(http.StatusOK, gin.H{
			"count": result["count"],
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"results": models.FilterResults(result["results"].([]bson.M)),
		})
	}
}

//处理批量删除 新增 修改等
func Batch(c *gin.Context)  {
	params := TranspilePostParams(c)
	fmt.Print("params is ", params)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
