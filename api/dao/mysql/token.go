package mysql

import (
	"api/model/entity"
	"log"
)

func InsertAccessToken(token *entity.AccessToken) error {
	sqlStr := "INSERT INTO token_table (id, user_id, expiration_time) VALUES (?,?,?)"
	insertStmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = insertStmt.Exec(token.ID, token.UserID, token.ExpirationTime)
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

func RetrieveAccessToken(tokenID string) *entity.AccessToken {
	sqlStr := "SELECT * FROM token_table WHERE id=?"
	queryRows := db.QueryRow(sqlStr, tokenID)
	accessToken := new(entity.AccessToken)
	err := queryRows.Scan(&accessToken.ID, &accessToken.UserID, &accessToken.ExpirationTime)
	if err != nil {
		log.Println(""+
			"Token db.QueryRow error:", err)
		return nil
	}
	return accessToken
}

func DeleteAccessToken(token *entity.AccessToken) error {
	sqlStr := "DELETE FROM token_table WHERE id=?"
	_, err := db.Exec(sqlStr, token.ID)
	if err != nil {
		log.Println("mysql delete token error:", err)
		return err
	}
	return nil
}
