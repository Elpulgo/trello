package commands

import (
	"fmt"
	"os"
	"trello/credentials"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tre",
	Short: "A simple CLI for Trello boards.",
	Long: `A simple CLI for Trello boards and tasks. 
	Built in Go for simplicity. Requires an API Key and Token for your Trello board.
	These will be stored in encrypted files on disk, protected by a password of choice.
	Optional to store the password on disk aswell if you don't want to be prompted on each use.
	Complete documentation is available at http://github.com/elpulgo/trello`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		credentialsExists, key, token := credentials.GetCredentials()

		if !credentialsExists {
			createCredentials()
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

func createCredentials() {
	var key string
	var token string
	var storePassphrase string
	var passphrase string

	fmt.Println(string("\033[32m"), "\n@ Credentials for Trello API not stored. Paste your Trello API key.")

	_, err := fmt.Scan(&key)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string("\033[32m"), "\n@ Paste your Trello API token.")

	_, err = fmt.Scan(&token)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string("\033[32m"), `
@ A password is required to access the Trello API credentials. 
  Would you like to store this password? (y/n)
  (Else you will be prompted each time for the password.)`)

	_, err = fmt.Scan(&storePassphrase)
	if err != nil {
		panic(err.Error())
	}

	if storePassphrase == "y" || storePassphrase == "Y" {
		fmt.Println(string("\033[32m"), "\n@ Enter passphrase to persist on disk (Will be saved in 'pass.dat')")
		_, err = fmt.Scan(&passphrase)
		if err != nil {
			panic(err.Error())
		}
		credentials.PersistPassphrase(passphrase)
	} else {
		fmt.Println(string("\033[32m"), "\n@ Enter passphrase to encrypt credentials. Note! This won't be saved so remember your passphrase!")
		_, err = fmt.Scan(&passphrase)
		if err != nil {
			panic(err.Error())
		}
	}

	credentials.PersistCredentials(
		key, token, passphrase)
}
