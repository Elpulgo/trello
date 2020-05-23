package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"trello/credentialsmanager"
	"trello/loader"
	"trello/models"

	color "trello/commandColors"

	"github.com/spf13/cobra"
)

var (
	specificBoard string
	boardName     string
	trelloKey     string
	trellotoken   string
	success       bool
)

var boardsCommand = &cobra.Command{
	Use:   "boards",
	Short: "Show all boards for user",
	Long:  `Show all boards for the user, in alphabetical order with numeric shortkey`,
	Run: func(cmd *cobra.Command, args []string) {
		loader.Run()
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()

		if !success {
			fmt.Println("Failed to get credentials. Bye bye.")
			return
		}

		if specificBoard == "" && boardName == "" {
			printBoards()
		} else {
			printCards()
		}
	},
}

func init() {
	boardsCommand.Flags().StringVarP(&specificBoard, "board", "b", "", "Show cards on a specific board, specified with either # or id.")
	boardsCommand.Flags().StringVarP(&boardName, "name", "n", "", "Show cards on a specific board, specified with a name.")

	rootCmd.AddCommand(boardsCommand)
}

func printBoards() {
	boards := GetAllBoards()

	loader.End()

	const padding = 16
	writer := tabwriter.NewWriter(os.Stdout, 10, 18, padding, '\t', tabwriter.AlignRight)
	defer writer.Flush()

	fmt.Fprintf(
		writer,
		"%s\t%s\t\t%s\t\n",
		color.GreenBold("Numeric short"),
		color.GreenBold("Name"),
		color.GreenBold("Id"))

	fmt.Fprintf(
		writer,
		"%s\t%s\t\t%s\t\n\n",
		color.GreenBold("-------------"),
		color.GreenBold("----"),
		color.GreenBold("--"))

	for index, board := range boards {
		fmt.Fprintf(
			writer,
			"[%s]\t\t%s\t{%s}\t\n",
			color.Yellow("# "+strconv.Itoa(index)),
			color.Cyan(board.Name),
			color.GreenBold(board.Id))
	}
}

func printCards() {
	var actions []models.Action
	var lists []models.List
	actionsChannel := make(chan []models.Action)
	listsChannel := make(chan []models.List)

	if boardName != "" {
		fmt.Println("Filter on board name")
	} else if len(specificBoard) > 2 {
		go GetCards(specificBoard, actionsChannel)
		actions = <-actionsChannel

		go GetLists(specificBoard, listsChannel)
		lists = <-listsChannel
	} else if index, err := strconv.Atoi(specificBoard); err == nil {
		boards := GetAllBoards()
		if index > len(boards) {
			loader.End()
			fmt.Println(color.RedBold("Short number is out of bounds. Check for boards and try again."))
			os.Exit(1)
		}

		fmt.Println("Need to collect all boards first and filter")
	} else {
		loader.End()
		fmt.Println(color.RedBold("Incorrect flags. Pass -h for help."))
		os.Exit(1)
	}

	var listMap []models.ListMap

	for _, m := range lists {
		listMap = append(listMap, models.ListMap{Id: m.Id, Name: m.Name, Actions: mapCardsToList(actions, m.Id)})
	}

	loader.End()

	for _, list := range listMap {
		fmt.Println("## " + list.Name)
		divider := strings.Repeat("=", len(list.Name)+3)
		fmt.Print(divider)
		fmt.Println("")

		for _, action := range list.Actions {
			if action.Name == "" {
				continue
			}

			commentsString := "     "
			if action.Badge.Comments > 0 {
				commentsString = "(*" + strconv.Itoa(action.Badge.Comments) + ")"
			}

			fmt.Println("{" + action.Id + "}  " + commentsString + "\t" + action.Name)
		}

		fmt.Println("")
	}
}

func mapCardsToList(allActions []models.Action, listId string) []models.Action {
	var cards []models.Action

	for _, action := range allActions {
		if action.ListId != listId {
			continue
		}

		cards = append(cards, action)
	}

	return cards
}
