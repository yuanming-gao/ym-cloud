//Package mysql :
// @Time : 2019/12/2 2:25 下午
// @Author : GaoYuanMing
// @Package : mysql
// @FileName : note.go
package mysql

import (
	"api/model/entity"
	"api/tools"
	"log"
)

func InsertNotes(notes *entity.Notes) error {
	sqlStr := "INSERT INTO notes_table (user_id, user_name, title, content, html, tags, create_time, update_time) VALUES (?,?,?,?,?,?,?,?)"
	insertStmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = insertStmt.Exec(notes.UserID,
		notes.UserName,
		notes.Title,
		notes.Content,
		notes.Html,
		notes.Tags,
		tools.NowTimeToUnixNano(),
		tools.NowTimeToUnixNano())
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err = insertStmt.Close(); err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func DeleteNotes(notesID int) error {
	sqlStr := "DELETE FROM notes_table WHERE id=?"
	_, err := db.Exec(sqlStr, notesID)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return err
	}
	return nil
}

func UpdateNotes(notes *entity.Notes) error {
	sqlStr := "UPDATE notes_table SET title=?, content=?, html=?,tags=?, update_time=? WHERE id=?"
	_, err := db.Exec(sqlStr, notes.Title, notes.Content, notes.Html, notes.Tags, tools.NowTimeToUnixNano(), notes.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetNotesInfoListByUserID(userId int) ([]*entity.Notes, error) {
	sqlStr := `select id, title,tags, create_time from notes_table where user_id=?`
	rows, err := db.Query(sqlStr, userId)
	defer func() {
		if rows != nil {
			if err = rows.Close(); err != nil {
				log.Println("Defer close rows error:", err)
			}
		}
	}()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	notesList := make([]*entity.Notes, 0)
	for rows.Next() {
		notes := new(entity.Notes)
		err := rows.Scan(
			&notes.ID,
			&notes.Title,
			&notes.Tags,
			&notes.CreateTime)
		if err != nil {
			log.Println("rows scan error:", err)
			return nil, err
		}
		notesList = append(notesList, notes)
	}
	return notesList, nil
}

func SelectNotesByID(id int) (*entity.Notes, error) {
	sqlStr := "SELECT * FROM notes_table WHERE id=?"
	queryRows := db.QueryRow(sqlStr, id)
	notes := new(entity.Notes)
	err := queryRows.Scan(&notes.ID, &notes.UserID, &notes.UserName, &notes.Title, &notes.Content, &notes.Html, &notes.Tags, &notes.CreateTime, &notes.UpdateTime)
	if err != nil {
		log.Println("db.QueryRow error:", err)
		return nil, err
	}
	return notes, nil
}
