package models

type Board struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Action struct {
	Id       string `json:"id"`
	ListId   string `json:"idList"`
	Name     string `json:"name"`
	Badge    Badge  `json:"badges"`
	Comments []Comment
}

type Badge struct {
	Comments int `json:"comments"`
}

type List struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"pos"`
}

type ListMap struct {
	Id      string
	Name    string
	Actions []Action
}

type Comment struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Text string `json:"text"`
	Card Card   `json:"card"`
}

type Card struct {
	Name string `json:"name"`
}
