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
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	// Read the keyboad input.
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		fmt.Println("Will exit")
	// 		os.Exit(1)
	// 	}

	// 	if input == "exit" {
	// 		fmt.Println("Will exit on users request")
	// 		os.Exit(1)
	// 	}
	commands.Execute()

	// }
	// commands.Execute()
}
