package commands

import (
	"fmt"
	"os"
	color "trello/commandColors"
	"trello/credentialsmanager"
	"trello/loader"

	"github.com/spf13/cobra"
)

var (
	listId      string
	title       string
	description string
)

var addCardCommand = &cobra.Command{
	Use:   "add-c",
	Short: "Add a new card to a specified list.",
	Long: `Add a new card to a specified list.
Requires listid and title.`,
	Run: func(cmd *cobra.Command, args []string) {
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()

		if !success {
			os.Exit(1)
		}

		loader.Run()

		if title == "" || listId == "" {
			loader.End()
			fmt.Println(color.RedBold("Title and list id are required and can't be empty."))
			os.Exit(1)
		}

		loader.End()
		fmt.Println("Hello add card! " + title + " - " + listId)
	},
}

func init() {
	addCardCommand.Flags().StringVarP(&title, "title", "t", "", "(*) Title of the card")
	addCardCommand.Flags().StringVarP(&listId, "listId", "l", "", "(*) List id. The list the card should belong to.")
	addCardCommand.Flags().StringVarP(&description, "description", "d", "", "Description for the card.")

	rootCmd.AddCommand(addCardCommand)
	credentialsCommand.MarkFlagRequired("title")
	credentialsCommand.MarkFlagRequired("listId")
}

// func addCardUrl(cardId string) string {
// 	return "https://api.trello.com/1/cards/" + cardId + "/actions?fields=type,data&key=" + trelloKey + "&token=" + trellotoken
// }
