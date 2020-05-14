package models

type Board struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Action struct {
	id   string     `json:"id`
	data actionData `json:"data"`
}

type actionData struct {
	text     string `json:"text"`
	cardId   string `json:"card:id"`
	listId   string `json:list:id`
	listName string `json:list:name`
}
