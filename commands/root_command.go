package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "2spiicrasher",
	Short: "2SpiiCrasher is a software allowing the connection of an amount of bot",
	Long:  `2SpiiCrasher is a software allowing the connection of a bot amount on a target IP`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
