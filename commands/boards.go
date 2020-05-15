package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
	"trello/credentialsmanager"
	"trello/models"

	"github.com/spf13/cobra"
)

var (
	specificBoard string
	boardName     string
)

var boardsCommand = &cobra.Command{
	Use:   "boards",
	Short: "Show all boards for user",
	Long:  `Show all boards for the user, in alphabetical order with numeric shortkey`,
	Run: func(cmd *cobra.Command, args []string) {
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

	const padding = 8
	writer := tabwriter.NewWriter(os.Stdout, 18, 8, padding, '\t', tabwriter.AlignRight)
	defer writer.Flush()

	fmt.Fprintf(writer, "%s\t%s\t%s\t\n", "Numeric short", "Name", "Id")
	fmt.Fprintf(writer, "%s\t%s\t%s\t\n\n", "=============", "====", "==")

	for index, board := range boards {
		fmt.Fprintf(writer, "[# %s]\t%s\t{%s}\t\n", strconv.Itoa(index), board.Name, board.Id)
	}
}

func printCards() {
	var actions []models.Action

	if len(specificBoard) > 2 {
		actions = getCards(specificBoard)
	} else {
		fmt.Println("Need to collect all boards first and filter")
		// boards := getAllBoards()
	}

	output := make(map[string][]models.Action)

	for _, m := range actions {
		if m.Data.List.Name == "" {
			continue
		}

		output[m.Data.List.Name] = append(output[m.Data.List.Name], m)
	}

	for list, actions := range output {
		fmt.Println(" - " + list)
		for _, action := range actions {
			fmt.Println("\t" + action.Data.Card.Name)
			if action.Data.Text != "" {
				fmt.Println("\t\t" + action.Data.Text)
			}
		}
		fmt.Println("")
	}
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

func getCards(boardId string) []models.Action {
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
	return actions
}

func getAllBoardsUrl() string {
	success, key, token := credentialsmanager.GetCredentials()

	if !success {
		os.Exit(1)
	}

	return "https://api.trello.com/1/members/me/boards?key=" + key + "&token=" + token
}

func getCardsUrl(boardId string) string {
	success, key, token := credentialsmanager.GetCredentials()

	if !success {
		os.Exit(1)
	}

	return "https://api.trello.com/1/boards/" + boardId + "/actions?key=" + key + "&token=" + token
}
