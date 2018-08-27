package repositories

import (
	"github.com/auth-core/app/models"
	"github.com/auth-core/app/types"
	"github.com/auth-core/database"
)

func GetClientByUsername(data *types.RequestGetToken) (*models.ClientAuth, error) {
	var clientAuth models.ClientAuth
	sqlQuery := `
					SELECT 
					c.id,
					c.name,
					c.username,
					c.password,
					c.description,
					ct.id as client_type_id,
					ct.name as client_type_name,
					c.created_by,
					c.updated_by,
					c.created_at,
					c.updated_at,
					c.deleted_at,
					c.row_status
					FROM clients c 
					LEFT JOIN client_types ct ON c.client_type_id = ct.id
					WHERE c.username = $1
				`
	db, rows, err := database.Select(sqlQuery, data.Username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}

	err = db.ScanRows(rows, &clientAuth)
	if err != nil {
		return nil, err
	}

	db.Close()
	return &clientAuth, nil
}

func GetClientToken(clientID int32) (*models.ClientAuthToken, error) {
	var clientAuthToken models.ClientAuthToken
	sqlQuery := `
					SELECT 
					ctt.id,
					ctt.token,
					ctt.created_by,
					ctt.updated_by,
					ctt.created_at,
					ctt.updated_at,
					ctt.deleted_at,
					ctt.row_status,
					c.id as client_id
					FROM client_tokens ctt
					LEFT JOIN clients c ON ctt.client_id = c.id
					WHERE c.id = $1 AND ctt.deleted_at IS NULL
				`
	db, rows, err := database.Select(sqlQuery, clientID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}

	err = db.ScanRows(rows, &clientAuthToken)
	if err != nil {
		return nil, err
	}

	db.Close()
	return &clientAuthToken, nil
}

func InsertClientToken(data *types.InsertClientToken) error {

	sqlQuery := "INSERT INTO client_tokens (client_id, token, expired_at, created_at, created_by) " +
		"VALUES (?,?,?,?,?)"
	db, err := database.Insert(sqlQuery,
		data.ClientID, data.Token, data.ExpiredAt, data.CreatedAt, data.CreatedBy)
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func UpdateClientToken(data *types.UpdateClientToken) error {
	sqlQuery := "UPDATE client_tokens SET deleted_at = $1, updated_by = $2, updated_at = $3 WHERE client_id = $4"
	db, err := database.Update(sqlQuery, data.DeletedAt, data.ClientID, data.UpdatedAt, data.ClientID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
