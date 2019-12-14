//Package notes
//@ Author: Gao YuanMing
//@ Data: 2019/12/11 9:20 下午
//@ Description:

package notes

import (
	"api/dao/mysql"
	"api/model/dto"
	"api/model/entity"
	"api/tools"
	"strings"
)

//CreateNotesService :
func CreateNotesService(userID int, notesDto *dto.CreateNotesDto) error {
	//处理tags,添加分隔符
	var tagsString string
	for _, tag := range notesDto.Tags {
		tagsString += "&+" + tag
	}
	notes := &entity.Notes{
		UserID:   userID,
		UserName: notesDto.UserName,
		Title:    notesDto.Title,
		Content:  notesDto.Content,
		Tags:     tagsString,
	}
	return mysql.InsertNotes(notes)
}

//GetOneNotesByID :
func GetOneNotesByID(notesID int) (*dto.NotesDto, error) {
	n, err := mysql.SelectNotesByID(notesID)
	if err != nil {
		return nil, err
	}
	//解析tags
	tagsArray := strings.Split(n.Tags, "&+")
	tagsArray = append(tagsArray[1:])
	notes := &dto.NotesDto{
		ID:         n.ID,
		Title:      n.Title,
		Content:    n.Content,
		Tags:       tagsArray,
		CreateTime: tools.ParseUnixNanoToString(n.CreateTime),
		UpdateTime: tools.ParseUnixNanoToString(n.UpdateTime),
	}
	return notes, nil
}

//GetNotesListService :
func GetNotesListService(userID int) ([]*dto.NotesInfoDto, error) {
	list, err := mysql.GetNotesInfoListByUserID(userID)
	if err != nil {
		return nil, err
	}
	infoList := make([]*dto.NotesInfoDto, 0)
	for _, v := range list {

		//解析tags
		tagsArray := strings.Split(v.Tags, "&+")
		tagsArray = append(tagsArray[1:])

		info := &dto.NotesInfoDto{
			ID:         v.ID,
			Title:      v.Title,
			Tags:       tagsArray,
			CreateTime: tools.ParseUnixNanoToString(v.CreateTime),
		}
		infoList = append(infoList, info)
	}
	return infoList, nil
}

//UserEditNotesService :
func UserEditNotesService(userID int, notesID int, notesDto *dto.CreateNotesDto) error {
	newNotes := &entity.Notes{
		ID:       notesID,
		UserID:   userID,
		UserName: notesDto.UserName,
		Title:    notesDto.Title,
		Content:  notesDto.Content,
	}
	return mysql.UpdateNotes(newNotes)
}

//DeleteNotesService :
func DeleteNotesService(notesID int) error {
	return mysql.DeleteNotes(notesID)
}
