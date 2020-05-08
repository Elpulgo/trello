package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listsCommand = &cobra.Command{
	Use:   "lists",
	Short: "Show all lists for a board",
	Long:  `Show all lists for a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List command. Yippikajay!")
	},
}

func init() {
	rootCmd.AddCommand(listsCommand)
}
