package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	color "trello/commandColors"
	"trello/loader"
	"trello/models"
)

func GetAllBoards() []models.Board {
	response, error := http.Get(getAllBoardsUrl())
	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to get boards from Trello API. Bye bye."))
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to parse boards from Trello API response. Bye bye."))
		os.Exit(1)
	}

	var boards []models.Board
	json.Unmarshal(body, &boards)
	return boards
}

func GetCards(boardId string, result chan []models.Action) {
	response, error := http.Get(getCardsUrl(boardId))

	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to get cards from Trello API. Bye bye."))
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to parse cards from Trello API response. Bye bye."))
		os.Exit(1)
	}

	var actions []models.Action
	json.Unmarshal(body, &actions)

	result <- actions
}

func GetLists(boardId string, filter string, result chan []models.List) {
	response, error := http.Get(getListsUrl(boardId))

	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to get cards from Trello API. Bye bye."))
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(color.RedBold("\n@ Failed to parse cards from Trello API response. Bye bye."))
		os.Exit(1)
	}

	var lists []models.List
	json.Unmarshal(body, &lists)

	if filter != "" {
		lists = filterLists(lists, filter)
	}

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Position < lists[j].Position
	})

	result <- lists
}

func AddCard(title string, listId string, description string) {
	payload := map[string]string{"name": title, "idList": listId, "desc": description}
	jsonPayload, _ := json.Marshal(payload)

	response, error := http.Post(
		getAddCardUrl(),
		"application/json",
		bytes.NewBuffer(jsonPayload))

	if error != nil {
		fmt.Println(color.RedBold("\nFailed to add cards from Trello API. Bye bye."))
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(color.RedBold("\nFailed to parse response from Trello API response. Bye bye."))
		os.Exit(1)
	}

	var card models.Card
	if err := json.Unmarshal(body, &card); err != nil {
		fmt.Println(color.RedBold("\nSuccessfully added card but failed to read response. Bye bye."))
		os.Exit(1)
	}

	fmt.Println(color.GreenBold("Successfully added card with id: ") + "{" + color.Yellow(card.Id) + "}")
}

func filterLists(values []models.List, value string) []models.List {
	for _, list := range values {
		if strings.ToLower(list.Name) == strings.ToLower(value) {
			return []models.List{list}
		}
	}

	loader.End()
	fmt.Println(color.YellowBold("Didn't find list '" + value + "'. Bye bye."))
	os.Exit(1)
	return []models.List{}
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

func getAddCardUrl() string {
	return "https://api.trello.com/1/cards?pos=bottom" + "&key=" + trelloKey + "&token=" + trellotoken
}
