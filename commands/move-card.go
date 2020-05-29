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
	newListId  string
	moveCardId string
)

var moveCardCommand = &cobra.Command{
	Use:   "move",
	Short: "Move card to new list.",
	Long: `Move card to new list.
Requires listid and cardid.`,
	Run: func(cmd *cobra.Command, args []string) {
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()

		if !success {
			os.Exit(1)
		}

		loader.Run()

		if newListId == "" || moveCardId == "" {
			loader.End()
			fmt.Println(color.RedBold("List id and card id are required and can't be empty. Bye bye."))
			os.Exit(1)
		}

		loader.End()

		MoveCard(moveCardId, newListId)
		fmt.Println(color.GreenBold("Successfully moved card!"))
	},
}

func init() {
	moveCardCommand.Flags().StringVarP(&moveCardId, "cardId", "c", "", "(*) Id of the card")
	moveCardCommand.Flags().StringVarP(&newListId, "listId", "l", "", "(*) List id. The list the card should move to.")

	cardCommand.AddCommand(moveCardCommand)
	credentialsCommand.MarkFlagRequired("moveCardId")
	credentialsCommand.MarkFlagRequired("newListId")
}
