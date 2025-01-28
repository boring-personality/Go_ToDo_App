package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	Long:  "Task will be permanantly deleted",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
