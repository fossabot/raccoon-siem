package migrator

import "C"
import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/core/migrator/assets"
	"github.com/tephrocactus/raccoon-siem/core/migrator/migration"
	"time"
)

type configuration struct {
	cockroachHost string
	cockroachPort string
	cockroachScheme string
}

var (
	Cmd = &cobra.Command{
		Use:  "migrate",
		Args: cobra.ExactArgs(0),
		RunE: Run,
	}
	cfg = new(configuration)
)

// Инициализируем конфигурацию
func init() {
}

// Запускаем миграцию
func Run(cmd *cobra.Command, _ []string) error {
	if err := parseFlags(cmd); err != nil {
		return err
	}

	processingBegan := time.Now()
	cockroachMigration := migration.CockroachMigration{}
	cockroachMigrationFiles := assets.GetMigrationFiles()
	appliedFiles, err := cockroachMigration.Run(cfg.cockroachHost, cfg.cockroachPort, cfg.cockroachScheme, cockroachMigrationFiles)
	if err != nil {
		return err
	}

	migration.ReportAppliedFiles(appliedFiles, processingBegan)
	return nil
}

func parseFlags(cmd *cobra.Command) error {
	var err error
	cfg.cockroachHost, err = cmd.Flags().GetString("db.host")
	if err != nil {
		return err
	}
	cfg.cockroachPort, err = cmd.Flags().GetString("db.port")
	if err != nil {
		return err
	}
	cfg.cockroachScheme, err = cmd.Flags().GetString("db.scheme")
	if err != nil {
		return err
	}
	return nil
}