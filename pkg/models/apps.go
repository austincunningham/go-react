package apps

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type(
	AppModel struct {
		db *sql.DB
	}
)
