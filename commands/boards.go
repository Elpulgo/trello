package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"trello/credentialsmanager"
	"trello/models"

	"github.com/gookit/color"
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
	boards := getAllBoards()

	const padding = 16
	writer := tabwriter.NewWriter(os.Stdout, 10, 18, padding, '\t', tabwriter.AlignRight)
	defer writer.Flush()

	red := color.FgRed.Render
	green := color.Style{color.FgGreen, color.OpBold}.Render
	yellow := color.FgYellow.Render

	fmt.Fprintf(writer, "%s\t%s\t\t%s\t\n", green("Numeric short", color.OpBold), green("Name"), green("Id"))
	fmt.Fprintf(writer, "%s\t%s\t\t%s\t\n\n", green("============="), green("===="), green("=="))

	for index, board := range boards {
		fmt.Fprintf(writer, "[%s]\t\t%s\t{%s}\t\n", red("# "+strconv.Itoa(index)), yellow(board.Name), green(board.Id))
	}
}

func printCards() {
	var actions []models.Action
	var lists []models.List
	actionsChannel := make(chan []models.Action)
	listsChannel := make(chan []models.List)

	if len(specificBoard) > 2 {
		go getCards(specificBoard, actionsChannel)
		actions = <-actionsChannel

		go getLists(specificBoard, listsChannel)
		lists = <-listsChannel
	} else {
		fmt.Println("Need to collect all boards first and filter")
		// TODO: boards := getAllBoards()
	}

	var listMap []models.ListMap

	for _, m := range lists {
		listMap = append(listMap, models.ListMap{Id: m.Id, Name: m.Name, Actions: getCardsForList(actions, m.Id)})
	}

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

func getCardsForList(allActions []models.Action, listId string) []models.Action {
	var cards []models.Action

	for _, action := range allActions {
		if action.ListId != listId {
			continue
		}

		cards = append(cards, action)
	}

	return cards
}

func getAllBoards() []models.Board {
	response, error := http.Get(getAllBoardsUrl())
	if error != nil {
		fmt.Println("\n@ Failed to get boards from Trello API. Will exit.")
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("\n@ Failed to parse boards from Trello API response. Will exit.")
		os.Exit(1)
	}

	var boards []models.Board
	json.Unmarshal(body, &boards)
	return boards
}

func getCards(boardId string, result chan []models.Action) {
	response, error := http.Get(getCardsUrl(boardId))

	if error != nil {
		fmt.Println("\n@ Failed to get cards from Trello API. Will exit.")
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("\n@ Failed to parse cards from Trello API response. Will exit.")
		os.Exit(1)
	}

	var actions []models.Action
	json.Unmarshal(body, &actions)

	result <- actions
}

func getLists(boardId string, result chan []models.List) {
	response, error := http.Get(getListsUrl(boardId))

	if error != nil {
		fmt.Println("\n@ Failed to get cards from Trello API. Will exit.")
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("\n@ Failed to parse cards from Trello API response. Will exit.")
		os.Exit(1)
	}

	var lists []models.List
	json.Unmarshal(body, &lists)

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Position < lists[j].Position
	})

	result <- lists
}

func getAllBoardsUrl() string {
	return "https://api.trello.com/1/members/me/boards?key=" + trelloKey + "&token=" + trellotoken
}

func getCardsUrl(boardId string) string {
	return "https://api.trello.com/1/boards/" + boardId + "/cards?key=" + trelloKey + "&token=" + trellotoken
}

func getCardUrl(cardId string) string {
	return "https://api.trello.com/1/cards/" + cardId + "?key=" + trelloKey + "&token=" + trellotoken
}

func getListsUrl(boardId string) string {
	return "https://api.trello.com/1/boards/" + boardId + "/lists?key=" + trelloKey + "&token=" + trellotoken
}
