package main

import (
	"fmt"
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
	ValidArgs: []string{"core", "collector", "correlator", "resources", "player"},
	RunE:      cobra.OnlyValidArgs,
}

var version string
var versionCmd = &cobra.Command{
	Use:  "version",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	// Sub commands
	rootCmd.AddCommand(
		versionCmd,
		core.Cmd,
		collector.Cmd,
		correlator.Cmd,
		player.Cmd,
		resources.Cmd)
}

func main() {
	rootCmd.Execute()
}
