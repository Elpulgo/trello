package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var boardsCommand = &cobra.Command{
	Use:   "boards",
	Short: "Show all boards for user",
	Long:  `Show all boards for the user, in alphabetical order with numeric shortkey`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Boards command. Yippikajay!")
	},
}

func init() {
	rootCmd.AddCommand(boardsCommand)
}
