package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func SqlConfig() *sqlx.DB {
	var Db *sqlx.DB
	host := "localhost"
	port := "3306"
	username := "go_journey"
	password := "go_journey"
	dbName := "go_journey"

	Db = sqlx.MustConnect("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbName)

	return Db
}
