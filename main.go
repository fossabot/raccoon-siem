package main

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/resource"
)

var rootCMD = &cobra.Command{
	Use:           "raccoon <component>",
	Short:         "Raccoon SIEM component launcher",
	Args:          cobra.ExactArgs(1),
	Run:           func(cmd *cobra.Command, args []string) {},
	SilenceErrors: true,
}

func init() {
	rootCMD.AddCommand(resource.Cmd)
}

func main() {
	rootCMD.Execute()
}
