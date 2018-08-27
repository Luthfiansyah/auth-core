package repositories

import (
	"fmt"

	"github.com/auth-core/app/models"
	"github.com/auth-core/database"
)

func GetCityAll() ([]models.City, error) {
	query :=
		`
			SELECT * FROM indonesia_capitals
    	`
	var city []models.City
	db, rows, err := database.Select(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		u := models.City{}
		db.ScanRows(rows, &u)
		city = append(city, u)
	}

	db.Close()
	return city, nil
}

func GetCityByCoordinate(latitude string, longitude string, radius string) ([]models.City, error) {
	fmt.Println("######", latitude, longitude, radius)
	query :=
		`
			select * from (
				SELECT  *,( 3959 * acos( cos( radians($1) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians($2) ) + sin( radians($1) ) * sin( radians( latitude ) ) ) ) AS distance 
				FROM indonesia_capitals
			) al
			where distance < $3
			ORDER BY distance
    	`
	var city []models.City
	db, rows, err := database.Select(query, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		u := models.City{}
		db.ScanRows(rows, &u)
		city = append(city, u)
	}

	db.Close()
	return city, nil
}

func DeleteByID(id string) error {

	sqlQuery := "DELETE FROM city WHERE id = ?"
	db, err := database.Delete(sqlQuery, id)
	if err != nil {
		return err
	}

	db.Close()
	return nil
}
