package main

import(	
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	"fmt"
	_"github.com/lib/pq"
)
// posgres vars
var db *sql.DB
const (
  dbhost = "localhost"
  dbport = "5432"
  dbuser = "postgres"
  dbpass = "password"
  dbname = "postgres"
)

//App struct
type App struct {
	ID        string   `json:"id,omitempty"`
	Appname   string   `json:"Appname,omitempty"`
	Disabled  bool     `json:"disabled,omitempty"`
	GlobalDisableMessage string `json:"globalDisableMessage,omitempty"`
	Versions  *Versions `json:"versions,omitempty"`
  }
  
//Versions struct
type Versions struct {
Version  string `json:"version,omitempty"`
Disabled bool   `json:"disabled,omitempty"`
DisableMessage string `json:"disableMessage,omitempty"`
}
var apps []App

func main(){
	dbConnect()

	router := echo.New()
	router.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "hello world!")
	})
	router.Logger.Fatal(router.Start(":8001"))
}

func dbConnect(){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)
	
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Posgres DB!")
}
