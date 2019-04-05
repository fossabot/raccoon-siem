package globals

import (
	"database/sql"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	DBConn  *sql.DB
	UDBConn sqlbuilder.Database
)

// TODO move
func NewUdbConnection(dbHost, dbPort, dbScheme string) error {
	var err error
	UDBConn, err = db.ConnectUdb(dbHost, dbPort, dbScheme)
	if err != nil {
		return err
	}

	return nil
}
