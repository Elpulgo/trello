package models

type Board struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Action struct {
	Id      string `json:"id"`
	ListId  string `json:"idList"`
	Name    string `json:"name"`
	Badge   Badge  `json:"badges"`
	Comment string
}

type Badge struct {
	Comments int64 `json:"comments"`
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
