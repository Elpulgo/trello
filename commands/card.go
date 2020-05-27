package commands

import (
	"fmt"
	"os"
	"strings"
	color "trello/commandColors"
	"trello/credentialsmanager"
	"trello/loader"
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

		loader.Run()

		if cardId == "" {
			loader.End()
			fmt.Println("Card id can't be empty. Specify a card id.")
			os.Exit(1)
		}

		printCard()
	},
}

func init() {
	cardCommand.Flags().StringVarP(&cardId, "id", "c", "", "(*) Card Id, required.")

	rootCmd.AddCommand(cardCommand)
	credentialsCommand.MarkFlagRequired("id")
}

func printCard() {
	var card models.Card
	var comments []models.Comment

	commentsChannel := make(chan []models.Comment)
	cardChannel := make(chan models.Card)

	go GetComments(cardId, commentsChannel)
	go GetCard(cardId, cardChannel)

	card = <-cardChannel
	comments = <-commentsChannel

	loader.End()

	fmt.Println("")
	fmt.Println(color.GreenBold(card.Name) + " {" + color.Yellow(card.Id) + "}")

	if card.Description != "" {
		fmt.Println(color.DarkGrey(card.Description))
	}

	divider := strings.Repeat(color.GreenBold("-"), len(card.Name)-2)
	fmt.Print(divider)
	fmt.Println("")

	for _, comment := range comments {
		if comment.Data.Text == "" {
			continue
		}

		commentText := strings.Replace(comment.Data.Text, "\n", "\n\t", -1)
		commentText = strings.Replace(commentText, "\r", "\n\t", -1)
		fmt.Println(
			"[" +
				color.Green(comment.Member.FullName) +
				color.DarkGreyBold(" @ ") +
				color.Yellow(comment.Date.Format("2006-01-02 15:12")) +
				"]\n\n" +
				color.Cyan(commentText))
		fmt.Println("")
	}
}
