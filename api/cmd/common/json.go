//Package handler :
/**
* @Time : 2019/12/1 2:40 下午
* @Author : GaoYuanMing
* @Package : handler
* @FileName : comm.go
 */
package common

import (
	jsonIter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary

func ParseAPIRequestJSON(request *http.Request, dto interface{}) error {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Println("defer func request body close error: ", err)
		}
	}()
	return json.Unmarshal(data, dto)
}
