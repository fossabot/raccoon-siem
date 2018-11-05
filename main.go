package main

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/collector"
	"github.com/tephrocactus/raccoon-siem/core"
	"github.com/tephrocactus/raccoon-siem/correlator"
	"github.com/tephrocactus/raccoon-siem/player"
	"github.com/tephrocactus/raccoon-siem/resources"
)

var rootCmd = &cobra.Command{
	Use:       "raccoon",
	Short:     "Raccoon SIEM component launcher",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"core", "collector", "correlator", "resources", "player", "generate"},
	RunE:      cobra.OnlyValidArgs,
}

func init() {
	// Sub commands
	rootCmd.AddCommand(
		core.Cmd,
		collector.Cmd,
		correlator.Cmd,
		player.Cmd,
		resources.Cmd)
}

func main() {
	rootCmd.Execute()
}
