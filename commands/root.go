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
	fmt.Println("== Credentials for Trello API not stored")

	var key string
	var token string
	var storePassphrase string
	var passphrase string

	fmt.Println(string("\033[32m"), "[] Paste your Trello API key")

	_, err := fmt.Scan(&key)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string("\033[32m"), "[] Paste your Trello API token")

	_, err = fmt.Scan(&token)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string("\033[32m"), "[] A password is required to access the Trello API credentials. Would you like to store this password? (y/n)\n(Else you will be prompted each time for the password.)")

	_, err = fmt.Scan(&storePassphrase)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string("\033[32m"), "[] Enter passphrase.")
	_, err = fmt.Scan(&passphrase)
	if err != nil {
		panic(err.Error())
	}

	if storePassphrase == "y" || storePassphrase == "Y" {
		trellohandler.PersistPassphrase(passphrase)
	}

	trellohandler.PersistCredentials(
		key, token, passphrase)
}
