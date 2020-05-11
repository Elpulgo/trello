package commands

import (
	"fmt"
	"os"
	"trello/trellohandler"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trello",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		credentialsExists, key, token := trellohandler.GetCredentials()

		if !credentialsExists {
			fmt.Println("Credentials don't exists")
		}

		fmt.Println("Key:" + key)
		fmt.Println("Token:" + token)

		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
