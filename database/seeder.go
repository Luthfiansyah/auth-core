package database

import "github.com/jinzhu/gorm"

func TableSeed(db *gorm.DB) error {
	// SeedClientType(db)
	// SeedClient(db)
	return nil
}

func SeedClientType(db *gorm.DB) error {

	var query string = `
		insert into client_types (name, description, created_by, created_at)
		values ('partner', 'partner', 0, '2018-02-05 21:10:58');
	`
	db.Exec(query)

	return nil
}

func SeedClient(db *gorm.DB) error {

	var query string = `
		insert into clients (name, client_type_id, username, password, description, created_by, created_at)
		values ('admin', 1,'admin', '$2y$12$BbpykGUOnwJIpX1m28iRGutEbwyb84ZPDX5sSYl1J7HvkdGpG0Jea', 'admin', 0, '2018-02-05 21:10:58');
	`
	db.Exec(query)

	return nil
}
