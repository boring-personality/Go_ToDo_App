package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to list",
	Long:  "Task will be added to the task list and marked as incomplete",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		add(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
