package generate

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/resources/generate/active_list"
	"github.com/tephrocactus/raccoon-siem/resources/generate/elastic"
)

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "generate content",
	Args:  cobra.ExactArgs(1),
}

func init() {
	Cmd.AddCommand(
		elastic.Cmd,
		activeList.Cmd,
	)
}
