package apps

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type(
	AppModel struct {
		db *sql.DB
	}
    //App struct
	App struct {
		ID                   string    `json:"id,omitempty"`
		Appname              string    `json:"Appname,omitempty"`
		Disabled             bool      `json:"disabled,omitempty"`
		GlobalDisableMessage string    `json:"globalDisableMessage,omitempty"`
		Versions             *Versions `json:"versions,omitempty"`
	}
	//Versions struct
	Versions struct {
		Version        string `json:"version,omitempty"`
		Disabled       bool   `json:"disabled,omitempty"`
		DisableMessage string `json:"disableMessage,omitempty"`
	}
)
// //App struct
// type App struct {
// 	ID                   string    `json:"id,omitempty"`
// 	Appname              string    `json:"Appname,omitempty"`
// 	Disabled             bool      `json:"disabled,omitempty"`
// 	GlobalDisableMessage string    `json:"globalDisableMessage,omitempty"`
// 	Versions             *Versions `json:"versions,omitempty"`
// }

// //Versions struct
// type Versions struct {
// 	Version        string `json:"version,omitempty"`
// 	Disabled       bool   `json:"disabled,omitempty"`
// 	DisableMessage string `json:"disableMessage,omitempty"`
// }

// var apps []App