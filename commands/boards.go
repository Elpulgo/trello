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
	listName      string
	trelloKey     string
	trellotoken   string
	success       bool
)

var boardsCommand = &cobra.Command{
	Use:   "boards",
	Short: "Show all boards for user",
	Long:  `Show all boards for the user, with numeric shortkey, name and id.`,
	Run: func(cmd *cobra.Command, args []string) {
		success, trelloKey, trellotoken = credentialsmanager.GetCredentials()
		loader.Run()

		if !success {
			loader.End()
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
	boardsCommand.Flags().StringVarP(&listName, "listname", "l", "", "Pass listname, for a specific list. Must be combined with -b/-n")

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
			"[%s]\t\t%s\t{%s}\t{%s}\n",
			color.Yellow("# "+strconv.Itoa(index)),
			color.Cyan(board.Name),
			color.DarkGreyBold(board.Id),
			color.DarkGrey(board.ShortUrl))
	}
}

func printCards() {
	var actions []models.Action
	var lists []models.List
	actionsChannel := make(chan []models.Action)
	listsChannel := make(chan []models.List)

	if boardName != "" {
		boards := GetAllBoards()
		success, board, suggestions := findBoardByFreeText(boardName, boards)
		if !success {
			loader.End()
			if len(suggestions) > 0 {
				fmt.Println(color.YellowBold("Failed to find your board, did you mean?\n"))
				for _, suggestion := range suggestions {
					fmt.Println("* " + suggestion)
				}
			} else {
				fmt.Println(color.YellowBold("Failed to find your board!\n"))
			}
			os.Exit(1)
		}

		specificBoard = board.Id
	} else if len(specificBoard) <= 2 {
		if index, err := strconv.Atoi(specificBoard); err == nil {
			boards := GetAllBoards()
			if index > len(boards)-1 {
				loader.End()
				fmt.Println(color.RedBold("Numeric short is out of bounds. Check for boards and try again."))
				os.Exit(1)
			}
			specificBoard = boards[index].Id
		} else {
			loader.End()
			fmt.Println(color.RedBold("Numeric short can't be parsed correctly. Bye bye."))
			os.Exit(1)
		}
	}

	go GetCards(specificBoard, actionsChannel)
	actions = <-actionsChannel

	go GetLists(specificBoard, listName, listsChannel)
	lists = <-listsChannel

	if len(lists) < 1 {
		loader.End()
		fmt.Println(color.RedBold("No lists exists on choosen board. Bye bye."))
		os.Exit(1)
	}

	var listMap []models.ListMap

	for _, m := range lists {
		listMap = append(listMap, models.ListMap{Id: m.Id, Name: m.Name, Actions: mapCardsToList(actions, m.Id)})
	}

	loader.End()
	fmt.Println("")
	for _, list := range listMap {
		fmt.Println(color.GreenBold(list.Name))
		fmt.Println("{" + color.DarkGrey(list.Id) + "}")
		divider := strings.Repeat(color.GreenBold("-"), len(list.Name))
		fmt.Print(divider)
		fmt.Println("")

		for _, action := range list.Actions {
			if action.Name == "" {
				continue
			}

			commentsString := ""
			if action.Badge.Comments > 0 {
				commentsString = "[C* " + strconv.Itoa(action.Badge.Comments) + "]"
			} else {
				commentsString = "        "
			}

			if action.Badge.Attachments > 0 {
				commentsString = commentsString + " (A* " + strconv.Itoa(action.Badge.Attachments) + ")"
			} else {
				commentsString = commentsString + "        "
			}

			actionName := strings.Replace(action.Name, "\n", "\t\t", -1)
			fmt.Println("{" + color.Yellow(action.Id) + "}  " + commentsString + "\t" + color.Cyan(actionName))
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

func findBoardByFreeText(freeText string, boards []models.Board) (bool, models.Board, []string) {
	var suggestions []string

	freeText = strings.ToLower(freeText)

	for _, board := range boards {
		if strings.ToLower(board.Name) == freeText {
			return true, board, []string{}
		}

		if strings.Index(strings.ToLower(board.Name), freeText) > -1 {
			suggestions = append(suggestions, board.Name)
		}
	}

	return false, models.Board{}, suggestions
}
