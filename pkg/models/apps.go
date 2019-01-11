package apps

import (
	"database/sql"
)

type(
	// AppModel not sure it should be called this
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
