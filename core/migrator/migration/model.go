package migration

type migrationModel struct {
	Version     string `db:"version" json:"version"`
	Description string `db:"description" json:"description"`
	Content     string `db:"content" json:"content"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
}
