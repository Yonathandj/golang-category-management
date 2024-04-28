package config

import (
	"database/sql"
	"fmt"
	"golang-category-management/helper"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "enigma02sukses"
	dbname   = "category_management_db"
)

func DatabaseConnection() *sql.DB {
	var dataSource = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSource)
	helper.HelperPanic(err)

	err = db.Ping()
	helper.HelperPanic(err)

	return db
}
