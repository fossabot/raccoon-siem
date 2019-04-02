package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(host, port, schema string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", defaultUser, host, port, schema)
	return sql.Open("postgres", connectionString)
}
