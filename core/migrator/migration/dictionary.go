package migration

import "regexp"

var (
	migrationsTable         = "_migrations"
	cockroachMigrationRegex = regexp.MustCompile("(V\\d+)__(.+)\\.sql")
)
