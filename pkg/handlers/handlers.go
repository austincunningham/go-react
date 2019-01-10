package handlers

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/austincunningham/go-react/pkg/db"
)

var d = db.DBconnect()
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
	rows, err := d.Query(sqlStatment)
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

// // GetApp gets an app by id
func GetApp(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("id passed in : ", id)
	var app App
	sqlStatment := `SELECT id, appname, disabled, globaldisablemessage FROM apps WHERE id=$1;`
	row := d.QueryRow(sqlStatment, id)
	err := row.Scan(&app.ID, &app.Appname, &app.Disabled, &app.GlobalDisableMessage)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, app)
}

// UpdateApp based on id app/:id, e.g. object
// {
//  "id":"1"
// 	"Appname": "Minishift-RHMAP",
// 	"Disabled": true,
// 	"globalDisableMessage": "disabled by API insomnia"
// }
func UpdateApp(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("id passed in : ", id)
	var app = new(App)
	if err := c.Bind(app); err != nil {
		fmt.Println(err)
		return err
	}
	sqlStatment := "UPDATE apps SET appname=$1, disabled=$2, globaldisablemessage=$3 WHERE id=$4"
	res, err := d.Query(sqlStatment, app.Appname, app.Disabled, app.GlobalDisableMessage, app.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, app)
	}

	return c.JSON(http.StatusOK, app.ID)
}

// CreateApp post json to create app in postgres db e.g. object
// {
// 	"Appname": "Minishift-RHMAP",
// 	"Disabled": true,
// 	"globalDisableMessage": "disabled by API insomnia"
// }
func CreateApp(c echo.Context) error {
	var app = new(App)
	if err := c.Bind(app); err != nil {
		fmt.Println(err)
		return err
	}
	sqlStatment := "INSERT INTO apps (appname, disabled, globaldisablemessage)VALUES ($1,$2,$3)"
	res, err :=d.Query(sqlStatment,app.Appname, app.Disabled, app.GlobalDisableMessage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, app)
	}

	return c.JSON(http.StatusOK, app.ID)

}
