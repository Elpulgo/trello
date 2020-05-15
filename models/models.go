package models

type Board struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Action struct {
	Id   string     `json:"id`
	Data actionData `json:"data"`
	Type string     `json:"type"`
}

type actionData struct {
	Text string `json:"text"`
	Card card   `json:"card"`
	List list   `json:list`
}

type card struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type list struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
