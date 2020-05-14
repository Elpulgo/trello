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

var boardsCommand = &cobra.Command{
	Use:   "boards",
	Short: "Show all boards for user",
	Long:  `Show all boards for the user, in alphabetical order with numeric shortkey`,
	Run: func(cmd *cobra.Command, args []string) {
		printBoards()
	},
}

func init() {
	rootCmd.AddCommand(boardsCommand)
}

func printBoards() {
	response, error := http.Get(getUrl())
	if error != nil {
		fmt.Println("\n@ Failed to get boards from Trello API. Will exit.")
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("\n@ Failed to parse boards from Trelo API response. Will exit.")
		os.Exit(1)
	}

	var boards []models.Board
	json.Unmarshal(body, &boards)

	const padding = 8
	writer := tabwriter.NewWriter(os.Stdout, 18, 8, padding, '\t', tabwriter.AlignRight)
	defer writer.Flush()

	fmt.Fprintf(writer, "%s\t%s\t%s\t\n", "Numeric short", "Name", "Id")
	fmt.Fprintf(writer, "%s\t%s\t%s\t\n\n", "=============", "====", "==")

	for index, board := range boards {
		fmt.Fprintf(writer, "[# %s]\t%s\t{%s}\t\n", strconv.Itoa(index), board.Name, board.Id)
	}
}

func getUrl() string {
	success, key, token := credentialsmanager.GetCredentials()

	if !success {
		os.Exit(1)
	}

	return "https://api.trello.com/1/members/me/boards?key=" + key + "&token=" + token
}
