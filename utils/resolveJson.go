package utils
import (
	"io/ioutil"
	"encoding/json"
	"fmt"
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