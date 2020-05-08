package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tasksCommand = &cobra.Command{
	Use:   "tasks",
	Short: "Show all tasks for a board",
	Long:  `Show all tasks for a specific board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List command. Yippikajay!")
	},
}

func init() {
	rootCmd.AddCommand(tasksCommand)
}
