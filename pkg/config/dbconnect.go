package config

import (
	"database/sql"
	"fmt"
)
// posgres vars should be using env vars for this
var DB *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "password"
	dbname = "postgres"
)


// DbConnect Connect to postgres database
func DbConnect() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		//using panic may not be best practice
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Posgres DB!")
}
