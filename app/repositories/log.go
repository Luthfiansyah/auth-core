package repositories

import (
	"github.com/auth-core/app/models"
	"github.com/auth-core/database"
)

func InsertLog(data *models.Log) error {

	sqlQuery := "INSERT INTO logs " +
		"( endpoint, request_message, response_message, request_time, response_time, elapsed_time,  " +
		"created_at, created_by) " +
		"VALUES (?, ?,?,?,?,?,?,?)"
	db, err := database.Insert(sqlQuery,
		data.Endpoint, data.RequestMessage, data.ResponseMessage, data.RequestTime, data.ResponseTime, data.ElapsedTime, data.CreatedAt, data.CreatedBy)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
