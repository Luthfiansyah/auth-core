package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/auth-core/app/models"
	"github.com/auth-core/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (err error) {
	db, err := DBOpen()
	migration(db)
	return err
}

func DBOpen() (db *gorm.DB, err error) {

	serverMode := config.MustGetString("server.mode")
	dbDriver := config.MustGetString(serverMode + ".db_driver")
	dbHost := config.MustGetString(serverMode + ".db_host")
	dbPort := config.MustGetString(serverMode + ".db_port")
	dbName := config.MustGetString(serverMode + ".db_name")
	dbUsername := config.MustGetString(serverMode + ".db_username")
	dbPassword := config.MustGetString(serverMode + ".db_password")

	if dbDriver == "postgres" {
		connStringPostgres := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUsername, dbPassword, dbName, dbHost, dbPort, "disable")
		db, err = gorm.Open("postgres", connStringPostgres)
	} else if dbDriver == "mysql" {
		connStringMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, err = gorm.Open("mysql", connStringMysql)
	} else {
		fmt.Println("No Database Selected!, Please check config.toml")
	}

	// defer db.Close()
	return db, err
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func migration(db *gorm.DB) {

	// TABLE LOG
	db.AutoMigrate(&models.Log{})

	// TABLE CLIENT TYPE
	db.AutoMigrate(&models.ClientType{})

	// TABLE CLIENT
	db.AutoMigrate(&models.Client{})

	// TABLE CLIENT TOKEN
	db.AutoMigrate(&models.ClientToken{})

	// DB SEEDER
	// db.Exec(seeder)
	TableSeed(db)

	return
}

func Select(query string, args ...interface{}) (*gorm.DB, *sql.Rows, error) {
	DB, err := DBOpen()
	if err != nil {
		return DB, nil, err
	}
	res, err := DB.Raw(query, args...).Rows()
	return DB, res, err
}

func Insert(query string, args ...interface{}) (*gorm.DB, error) {
	DB, err := DBOpen()
	if err != nil {
		return DB, err
	}
	tx := DB.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return DB, err
	}
	tx.Commit()
	return DB, nil
}

func Update(query string, args ...interface{}) (*gorm.DB, error) {
	DB, err := DBOpen()
	if err != nil {
		return DB, err
	}
	tx := DB.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return DB, err
	}
	tx.Commit()
	return DB, nil
}

func Delete(query string, args ...interface{}) (*gorm.DB, error) {
	DB, err := DBOpen()
	if err != nil {
		return DB, err
	}
	DB.Exec(query, args...)
	if err != nil {
		return DB, err
	}
	return DB, nil
}
