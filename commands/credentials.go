package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	key             string
	token           string
	passphrase      string
	storepassphrase string
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
		fmt.Println("Credentials command FIRED!" + key + " token: " + token + " pp:" + passphrase)
	},
}

func init() {

	credentialsCommand.Flags().StringVarP(&key, "trello-key", "k", "", "(*) Trello API key.")
	credentialsCommand.Flags().StringVarP(&token, "trello-token", "t", "", "(*) Trello API token.")
	credentialsCommand.Flags().StringVarP(&passphrase, "passphrase", "p", "", "(*) Passphrase for API credentials.")
	credentialsCommand.Flags().StringVarP(&storepassphrase, "store", "s", "", "Should store passphrase in 'pass.dat' (y/n)")

	rootCmd.AddCommand(credentialsCommand)

	credentialsCommand.MarkFlagRequired("trello-key")
	credentialsCommand.MarkFlagRequired("trello-token")
	credentialsCommand.MarkFlagRequired("passphrase")
}
