package handler

import (
	"api/cmd/common"
	"api/model/dto"
	notesService "api/service/notes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateNotes :创建笔记
func CreateNotes(c *gin.Context) {
	notesDto := new(dto.CreateNotesDto)
	responseBody := new(common.APIResponseBody)
	err := common.ParseAPIRequestJSON(c.Request, notesDto)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "api请求参数缺失"
		common.SendAPIResponse(c, responseBody)
		return
	}
	err = notesService.CreateNotesService(userID, notesDto)
	if err != nil {
		responseBody.Status = http.StatusInternalServerError
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "笔记上传成功"
	common.SendAPIResponse(c, responseBody)
}

//SelectNotes :根据id查找出笔记
func SelectNotes(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	notesID, err := strconv.Atoi(c.Params.ByName("notes_id"))
	if err != nil {
		responseBody.Msg = "api请求的参数缺失"
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	notes, err := notesService.GetOneNotesByID(notesID)
	if err != nil {
		responseBody.Msg = err.Error()
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "操作成功"
	responseBody.Data = make(map[string]interface{})
	responseBody.Data["notes"] = notes
	common.SendAPIResponse(c, responseBody)
}

//UpdateNodes :更新笔记
func UpdateNodes(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	notesID, err := strconv.Atoi(c.Params.ByName("notes_id"))
	if err != nil {
		responseBody.Msg = "api请求参数缺失"
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	notesDto := new(dto.CreateNotesDto)
	err = common.ParseAPIRequestJSON(c.Request, notesDto)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "请求的json数据错误"
		common.SendAPIResponse(c, responseBody)
		return
	}
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "api请求参数缺失" + err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	err = notesService.UserEditNotesService(userID, notesID, notesDto)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "操作成功"
	common.SendAPIResponse(c, responseBody)
}

//DeleteNotes 删除笔记
func DeleteNotes(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	notesID, err := strconv.Atoi(c.Params.ByName("notes_id"))
	if err != nil {
		responseBody.Msg = "api请求的参数缺失" + err.Error()
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	err = notesService.DeleteNotesService(notesID)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "操作成功"
	common.SendAPIResponse(c, responseBody)
}

//ListAllNotes :列出所有的笔记
func ListAllNotes(c *gin.Context) {
	responseBody := new(common.APIResponseBody)
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		responseBody.Msg = "api请求参数缺失" + err.Error()
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	notesList, err := notesService.GetNotesListService(userID)
	if err != nil {
		responseBody.Msg = err.Error()
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "查询成功"
	responseBody.Data = make(map[string]interface{})
	responseBody.Data["notes_number"] = len(notesList)
	responseBody.Data["notes_list"] = notesList
	common.SendAPIResponse(c, responseBody)
}
