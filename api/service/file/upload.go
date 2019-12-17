package file

import (
	"api/dao/mysql"
	"api/model/entity"
	"log"
)

func SaveUploadFileInfo(info *entity.File) error {
	return mysql.InsertFileInfo(info)
}

func Upload(objectName, localFileName string) error {
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
