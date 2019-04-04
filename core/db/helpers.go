package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

type QueryConfig struct {
	Tx sqlbuilder.Tx
	OrderBy []interface{}
	Page uint
}

func Connect(host, port, schema string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", defaultUser, host, port, schema)
	return sql.Open("postgres", connectionString)
}

func ConnectUdb(host, port, schema string) (sqlbuilder.Database, error) {
	return postgresql.Open(postgresql.ConnectionURL{
		User:     defaultUser,
		Host:     host,
		Database: schema,
		Options:  map[string]string{"port": port, "sslmode": "disable"},
	})
}

func IDEmpty(id string) bool {
	return id == "" || id == defaultUUID
}

