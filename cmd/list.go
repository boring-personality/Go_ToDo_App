package cmd

import (
	"github.com/spf13/cobra"
)

var All bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks",
	Long:  "incomplete tasks will be displayed",
	Args:  cobra.MatchAll(cobra.NoArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if All {
			listall()
		} else {
			list()
		}
	},
}

func init() {
	listCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "list all the completed and incomplete tasks")
	rootCmd.AddCommand(listCmd)
}
