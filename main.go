package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ToDo App")
	if len(os.Args) <= 1 {
		fmt.Println("No command specified")
		// Print the proper message here like how to use the tool
	} else {
		cmd := os.Args[1]
		fmt.Println("Command: ", cmd)
		switch cmd {
		case "add":
			fmt.Println("Yet to implement")
		case "complete":
			fmt.Println("Yet to implement")
		case "delete":
			fmt.Println("Yet to implement")
		case "list":
			list()
		case "help":
			fmt.Println("Yet to implement")
		default:
			fmt.Println("Invalid command")
		}
	}
}
