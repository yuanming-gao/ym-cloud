//Package mysql
//@ Author: Gao YuanMing
//@ Data: 2019/12/11 9:12 下午
//@ Description:

package mysql

import (
	"api/model/entity"
	"testing"
	"time"
)

func TestInsertNotes(t *testing.T) {
	ConnDatabaseAndCreateTables()
	notes := &entity.Notes{
		UserID:     1,
		UserName:   "高元明",
		Title:      "微服务设计",
		Content:    "近几年的时间啊骷简介三年级三季度看见你撒看见你",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := InsertNotes(notes)
	if err != nil {
		t.Error(err)
	}
}

func TestSelectNotesByID(t *testing.T) {
	ConnDatabaseAndCreateTables()
	notes := SelectNotesByID(1)
	t.Log(notes)
}

func TestUpdateNotes(t *testing.T) {
	ConnDatabaseAndCreateTables()
	notes := &entity.Notes{
		ID:         1,
		UserID:     1,
		UserName:   "高元明",
		Title:      "分布式系统",
		Content:    "kalkdkkaskldaklsklklakdkaskldakl",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := UpdateNotes(notes)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteNotes(t *testing.T) {
	ConnDatabaseAndCreateTables()
	notes := &entity.Notes{
		ID:         1,
		UserID:     1,
		UserName:   "高元明",
		Title:      "分布式系统",
		Content:    "kalkdkkaskldaklsklklakdkaskldakl",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := DeleteNotes(notes)
	if err != nil {
		t.Error(err)
	}
}
