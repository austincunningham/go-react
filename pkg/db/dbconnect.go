package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
// posgres vars should be using env vars for this
//var DB *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "password"
	dbname = "postgres"
)

type(
	AppModel struct {
		db *sql.DB
	}
)


// DbConnect Connect to postgres database
func DBconnect() *sql.DB{
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//using panic may not be best practice
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Posgres DB!")
	return db
}
