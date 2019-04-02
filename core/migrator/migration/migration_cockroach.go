package migration

import (
	"database/sql"
	"fmt"
	"github.com/jessevdk/go-assets"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"time"
)

type CockroachMigration struct{}

// Накатывает миграции. Содержимое каждого файла миграции накатывается в отдельной транзакции.
func (r *CockroachMigration) Run(host, port, schema string, migrationsToApply []*assets.File) ([]string, error) {
	// Подключаемся к дефольной схеме, чтобы создать указанную схему
	connection, err := db.Connect(host, port, "")
	if err != nil {
		return nil, err
	}

	// Создаем указанную схему, если еще не существует
	if err := r.createDatabase(connection, schema); err != nil {
		return nil, err
	}

	// Закрываем соединение, т.к. нет возможности динамически сменить активную схему.
	if err := connection.Close(); err != nil {
		return nil, err
	}

	// Открываем новое подключение.
	if connection, err = db.Connect(host, port, schema); err != nil {
		return nil, err
	}

	// На выходе закрываем последнее подключение.
	defer connection.Close()

	// Создаем таблицу миграций, если еще не существует
	if err := r.createMigrationsTable(connection); err != nil {
		return nil, err
	}

	// Читаем информацию о ранее примененных миграциях
	existingMigrations, err := r.readExistingMigrations(connection)
	if err != nil {
		return nil, err
	}

	// Список примененных файлов миграции (для возврата)
	var appliedFiles []string

	// Итерируемся по целевым миграциям
	for _, file := range migrationsToApply {
		// Валидируем и наполняем очередную миграцию метаданными из имени файла
		nextMigration, err := r.prepareMigration(file)
		if err != nil {
			return appliedFiles, err
		}

		// Если текущая миграция еще не была применена, применяем
		if !r.migrationAlreadyApplied(existingMigrations, nextMigration.Version) {
			if err := r.applyMigration(connection, nextMigration); err != nil {
				return appliedFiles, err
			}
			appliedFiles = append(appliedFiles, file.Name())
		}
	}

	return appliedFiles, nil
}

func (r *CockroachMigration) createDatabase(conn *sql.DB, schema string) error {
	createDatabaseQuery := fmt.Sprintf("create database if not exists %s;", schema)
	_, err := conn.Exec(createDatabaseQuery)
	return err
}

func (r *CockroachMigration) createMigrationsTable(conn *sql.DB) error {
	createTableQuery := fmt.Sprintf(`
		create table if not exists %s (
			version string primary key,
			description string not null,
			content string not null,
			created_at int not null
		);
	`, migrationsTable)
	_, err := conn.Exec(createTableQuery)
	return err
}

func (r *CockroachMigration) readExistingMigrations(conn *sql.DB) ([]*migrationModel, error) {
	var existingMigrations []*migrationModel
	selectMigrationsQuery := fmt.Sprintf(
		`SELECT version, description, content, created_at 
				FROM %s 
				ORDER BY version DESC;
	`, migrationsTable)
	rows, err := conn.Query(selectMigrationsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		migration := &migrationModel{}
		err := rows.Scan(&migration.Version, &migration.Description, &migration.Content, &migration.CreatedAt)
		if err != nil {
			return nil, err
		}
		existingMigrations = append(existingMigrations, migration)
	}

	return existingMigrations, nil
}

func (r *CockroachMigration) prepareMigration(file *assets.File) (*migrationModel, error) {
	match := cockroachMigrationRegex.FindStringSubmatch(file.Name())
	if match == nil {
		return nil, fmt.Errorf("migration file name must be: V<epoch>__<description>.sql")
	}
	return &migrationModel{
		Version:     match[1],
		Description: match[2],
		Content:     string(file.Data),
		CreatedAt:   time.Now().UnixNano(),
	}, nil
}

func (r *CockroachMigration) applyMigration(conn *sql.DB, migration *migrationModel) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(migration.Content); err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		INSERT INTO %s (version, description, content, created_at) 
		VALUES ($1, $2, $3, $4)
	`, migrationsTable),
		migration.Version,
		migration.Description,
		migration.Content,
		migration.CreatedAt,
	)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *CockroachMigration) migrationAlreadyApplied(existingMigrations []*migrationModel, nextMigrationVersion string) bool {
	for _, em := range existingMigrations {
		if em.Version == nextMigrationVersion {
			return true
		}
	}
	return false
}
