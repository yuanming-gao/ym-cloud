//Package handler
//@ Author: Gao YuanMing//@ Data: 2019/12/8 9:18 下午
//@ Description:

package handler

import (
	"api/cmd/common"
	"api/config"
	"api/model/entity"
	fileservice "api/service/file"
	"api/tools"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UploadFile :文件上传接口
func UploadFile(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "api请求参数缺失"
		common.SendAPIResponse(c, responseBody)
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("defer close file error:", err)
		}
	}()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}

	//暂存本地
	if err = ioutil.WriteFile(config.FileLocalPath+fileHeader.Filename, bytes, 0666); err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		log.Println(err)
		return
	}

	//上传阿里云oss
	sha := tools.NewSha256String(bytes)
	ossAddress := "file/" + fileHeader.Filename
	if err = fileservice.GetBucket().PutObject(ossAddress, file); err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		log.Println(err)
		return
	}

	fileInfo := &entity.File{
		Sha256Value: sha,
		Name:        fileHeader.Filename,
		Type:        fileHeader.Header.Get("Content-Type"),
		UserID:      userID,
		Size:        fileHeader.Size,
		LocalAt:     config.FileLocalPath + fileHeader.Filename,
		UploadAt:    ossAddress,
		Status:      0,
		CreateTime:  tools.NowTimeToUnixNano(),
	}

	//文件上传保存记录
	err = fileservice.SaveUploadFileInfo(fileInfo)
	if err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = "该文件已存在"
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "上传成功"
	common.SendAPIResponse(c, responseBody)
}

//DownloadFile :文件下载接口
func DownloadFile(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	//获取文件
	name := c.Param("name")
	if name == "" {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "请求参数错误"
		common.SendAPIResponse(c, responseBody)
		return
	}
	bytes, err := ioutil.ReadFile(config.FileLocalPath + name)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	c.Writer.Header().Add("content-disposition", "attachment; filename=\""+name+"\"")
	_, err = c.Writer.Write(bytes)
	if err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "下载成功"
}

func ListFile(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "api请求参数缺失"
		common.SendAPIResponse(c, responseBody)
		return
	}
	list, err := fileservice.GetUserFileList(userID)
	if err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "获取成功"
	responseBody.Data = make(map[string]interface{})
	responseBody.Data["file_list"] = list
	responseBody.Data["file_number"] = len(list)
	common.SendAPIResponse(c, responseBody)
}

func DeleteFile(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	sha := c.Param("sha")
	err := fileservice.DeleteFile(sha)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "删除文件请求参数缺失"
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "删除成功"
	common.SendAPIResponse(c, responseBody)
}

func OpenFile(c *gin.Context) {

}
