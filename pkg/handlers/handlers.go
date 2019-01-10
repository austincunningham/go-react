package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var db *sql.DB
//App struct
type App struct {
	ID                   string    `json:"id,omitempty"`
	Appname              string    `json:"Appname,omitempty"`
	Disabled             bool      `json:"disabled,omitempty"`
	GlobalDisableMessage string    `json:"globalDisableMessage,omitempty"`
	Versions             *Versions `json:"versions,omitempty"`
}

//Versions struct
type Versions struct {
	Version        string `json:"version,omitempty"`
	Disabled       bool   `json:"disabled,omitempty"`
	DisableMessage string `json:"disableMessage,omitempty"`
}

var apps []App

// GetAllApps gets all apps
func GetAllApps(c echo.Context) error {
	// returning static apps array
	sqlStatment := "SELECT id, appname, disabled, globaldisablemessage FROM apps order by id"
	rows, err := db.Query(sqlStatment)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// creates a new object from return from postgres
	var result []App
	for rows.Next() {
		var app App
		err2 := rows.Scan(&app.ID, &app.Appname, &app.Disabled, &app.GlobalDisableMessage)
		if err2 != nil {
			return err2
		}
		result = append(result, app)
	}
	// returns new object
	return c.JSON(http.StatusOK, result)
}
