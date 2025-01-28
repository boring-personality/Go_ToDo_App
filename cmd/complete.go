package cmd

import (
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "mark complete",
	Long:  "Task will be marked as complete",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		complete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
