package core

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/core/api"
	"github.com/tephrocactus/raccoon-siem/core/globals"
	"github.com/tephrocactus/raccoon-siem/core/migrator"
)

var (
	Cmd = &cobra.Command{
		Use:   "core",
		Short: "start raccoon configuration server",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	listen, dbHost, dbPort, dbScheme string
)

func init() {
	// Migration
	Cmd.AddCommand(migrator.Cmd)

	// Cockroach
	Cmd.PersistentFlags().StringVar(
		&dbHost,
		"db.host",
		"localhost",
		"database host",
	)
	Cmd.PersistentFlags().StringVar(
		&dbPort,
		"db.port",
		"26257",
		"database port",
	)
	Cmd.PersistentFlags().StringVarP(
		&dbScheme,
		"db.scheme",
		"d",
		"raccoon",
		"database scheme",
	)

	// Listen address
	Cmd.Flags().StringVarP(
		&listen,
		"listen",
		"l",
		":7220",
		"listen address")

}

func run(_ *cobra.Command, _ []string) error {
	// Open database
	if err := globals.NewUdbConnection(dbHost, dbPort, dbScheme); err != nil {
		return err
	}

	// Register http endpoints
	router := api.GetRouter()

	// Run http server
	return router.Run(listen)
}
