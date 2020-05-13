package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var credentialsCommand = &cobra.Command{
	Use:   "credentials",
	Short: "Set credentials for Trello API.",
	Long: `
Set credentials for Trello API.

================================= 

Token and key is required for use of the CLI.
Can be obtained @ https://trello.com/1/appKey/generate.
Protect the API credentials with a password of choice.
The password can be saved encoded in a file if you don't wan't to be prompted each time.
(Will be saved in 'pass.dat')
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Credentials command FIRED!")
	},
}

func init() {
	rootCmd.AddCommand(credentialsCommand)
}
