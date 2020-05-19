package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"trello/credentialsmanager"
	"trello/models"

	"github.com/spf13/cobra"
)

var (
	cardId string
)

var cardCommand = &cobra.Command{
	Use:   "card",
	Short: "Show specific info for card",
	Long:  `Show specific info for card`,
	Run: func(cmd *cobra.Command, args []string) {
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()

		if !success {
			os.Exit(1)
		}

		if cardId == "" {
			fmt.Println("Card id can't be empty. Specify a card id.")
			os.Exit(1)
		}

		printCard()
	},
}

func init() {
	cardCommand.Flags().StringVarP(&cardId, "id", "c", "", "Card Id, required.")
	rootCmd.AddCommand(cardCommand)
	credentialsCommand.MarkFlagRequired("id")
}

func printCard() {
	fmt.Println("Hello card!")
}

func getComments(cardId string, result chan []models.Comment) {
	response, error := http.Get(getActionUrl(cardId))

	if error != nil {
		fmt.Println("\n@ Failed to get card from Trello API. Will exit.")
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("\n@ Failed to parse card from Trello API response. Will exit.")
		os.Exit(1)
	}

	var comments []models.Comment
	json.Unmarshal(body, &comments)

	result <- comments
}

func getActionUrl(cardId string) string {
	return "https://api.trello.com/1/cards/" + cardId + "/actions?fields=type,data&key=" + trelloKey + "&token=" + trellotoken
}
