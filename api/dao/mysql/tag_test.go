//Package mysql
//@ Author: Gao YuanMing
//@ Data: 2019/12/11 9:41 下午
//@ Description:

package mysql

import (
	"api/model/entity"
	"testing"
)

func TestInsertTags(t *testing.T) {
	ConnDatabaseAndCreateTables()
	tag := new(entity.NotesTag)
	tag.NotesID = 1
	tag.Content = "分布式"
	err := InsertTags(tag)
	if err != nil {
		t.Error(err)
	}
}

func TestGetNotesAllTagsByNotesID(t *testing.T) {
	ConnDatabaseAndCreateTables()
	tags := GetNotesAllTagsByNotesID(1)
	for _, v := range tags {
		t.Log(v)
	}
	t.Log(len(tags))
}

func TestDeleteTagsByNotesID(t *testing.T) {
	ConnDatabaseAndCreateTables()
	err := DeleteTagsByNotesID(1)
	if err != nil {
		t.Error(err)
	}
}
