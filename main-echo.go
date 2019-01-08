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
	defer db.Close()
	populateArray()

	router := echo.New()
	router.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "hello world!")
	})
	 
	router.GET("/apps", GetAllApps)
	router.GET("/apps/:id", GetApp)

	router.Logger.Fatal(router.Start(":8001"))
}

// GetAllApps gets all apps
func GetAllApps(c echo.Context) error {
	return c.JSON(http.StatusOK, apps)
}

// GetApp gets an app by id
func GetApp(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("id passed in : ",id)
	return c.JSON(http.StatusOK, id)
}

// no real difference here
func dbConnect(){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)
	
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		//using panic may not be best practice
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Posgres DB!")
}

func populateArray(){
	//var err error
	//populate the array hard coded mock data at the moment 
	apps = append(apps, App{
	  ID: "1",
	  Appname: "MDC", 
	  Disabled: false, 
	  GlobalDisableMessage: "Disabled", 
	  Versions: &Versions{
		Version: "1.1.1", 
		Disabled: false}})
	apps = append(apps, App{
	  ID: "2", 
	  Appname: "Integreatly", 
	  Disabled: false,
	  GlobalDisableMessage: "Disabled", 
	  Versions: &Versions{
		Version: "1.0.1", 
		Disabled: false,
		DisableMessage: "Disabled by admin"}})
	apps = append(apps, App{
	  ID: "3", 
	  Appname: "RHMAP", 
	  Disabled: true,
	  GlobalDisableMessage: "Disabled", 
	  Versions: &Versions{
		Version: "4.6.2", 
		Disabled: true,
		DisableMessage: "Disabled by admin"}})
}