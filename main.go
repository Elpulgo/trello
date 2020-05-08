package main

import (
	"fmt"
	commands "trello/commands"
)

func init() {
	fmt.Println("Init trello")
}

func main() {
	fmt.Println("Main trello")
	commands.Execute()
}
