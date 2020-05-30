package models

import "time"

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
	Comments    int `json:"comments"`
	Attachments int `json:"attachments"`
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
	Type   string    `json:"type"`
	Data   Data      `json:"data"`
	Date   time.Time `json:"date"`
	Member Member    `json:"memberCreator"`
}

type Data struct {
	Text string `json:"text"`
	Card Card   `json:"card"`
}

type Member struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	FullName string `json:"fullName"`
}

type Card struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}
