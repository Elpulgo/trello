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
	comment       string
	commentCardId string
)

var addCommentCommand = &cobra.Command{
	Use:   "comment",
	Short: "Add a new comment to a specific card.",
	Long: `Add a new comment to a specific card.
Requires card id and coment.`,
	Run: func(cmd *cobra.Command, args []string) {
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()

		if !success {
			os.Exit(1)
		}

		loader.Run()

		if commentCardId == "" || comment == "" {
			loader.End()
			fmt.Println(color.RedBold("Comment and card id are required and can't be empty. Bye bye."))
			os.Exit(1)
		}

		loader.End()

		AddComment(comment, commentCardId)
	},
}

func init() {
	addCommentCommand.Flags().StringVarP(&comment, "comment", "c", "", "(*) Comment.")
	addCommentCommand.Flags().StringVarP(&commentCardId, "cardId", "i", "", "(*) Card id. The card the comment should belong to.")

	cardCommand.AddCommand(addCommentCommand)
	credentialsCommand.MarkFlagRequired("comment")
	credentialsCommand.MarkFlagRequired("commentCardId")
}
