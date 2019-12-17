//Package mysql :
// @Time : 2019/12/2 2:25 下午
// @Author : GaoYuanMing
// @Package : mysql
// @FileName : file.go
package mysql

import (
	"api/model/dto"
	"api/model/entity"
	"api/tools"
	"log"
)

func InsertFileInfo(info *entity.File) error {
	sqlStr := "INSERT INTO file_table (sha256_value, name, type,user_id,size,status,local_at,upload_at,create_time) VALUES (?,?,?,?,?,?,?,?,?)"
	insertStmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = insertStmt.Exec(info.Sha256Value,
		info.Name,
		info.Type,
		info.UserID,
		info.Size,
		info.Status,
		info.LocalAt,
		info.UploadAt,
		info.CreateTime)
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

func GetFileInfoListByUserId(userID int) ([]*dto.FileInfoDto, error) {
	sqlStr := `select sha256_value, create_time , name, local_at, size from file_table where user_id=?`
	rows, err := db.Query(sqlStr, userID)
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
	fileInfoArray := make([]*dto.FileInfoDto, 0)
	for rows.Next() {
		info := new(dto.FileInfoDto)
		var t, size int64
		err := rows.Scan(
			&info.Sha,
			&t,
			&info.Name,
			&info.Path,
			&size)
		if err != nil {
			log.Println("rows scan error:", err)
			return nil, err
		}
		info.Time = tools.ParseUnixNanoToString(t)
		info.Size = tools.ConversionBitSize(size)
		fileInfoArray = append(fileInfoArray, info)
	}
	return fileInfoArray, nil
}

func DeleteFileInfoBySha256Value(sha string) error {
	sqlStr := "DELETE FROM file_table WHERE sha256_value=?"
	_, err := db.Exec(sqlStr, sha)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return err
	}
	return nil
}
