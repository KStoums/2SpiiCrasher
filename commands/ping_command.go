package commands

import (
	"github.com/spf13/cobra"
)

var pingCommand = &cobra.Command{
	Use:     "ping",
	Short:   "Sends a defined ping count to the target server IP",
	Long:    "Sends a defined ping count to the target server IP",
	Example: "2spiicrasher ping 100 play.myserver.net",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// WORK
	},
}

func init() {
	rootCommand.AddCommand(pingCommand)
}
