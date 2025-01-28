package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ToDo App")
	params_len := len(os.Args)
	if params_len <= 1 {
		fmt.Println("No command specified")
		// Print the proper message here like how to use the tool
	} else {
		cmd := os.Args[1]
		var task string
		if params_len == 3 {
			task = os.Args[2]
		}
		fmt.Println("Command: ", cmd)
		switch cmd {
		case "add":
			add(task)
		case "complete":
			complete(task)
		case "delete":
			delete(task)
		case "list":
			list()
		case "listall":
			listall()
		case "help":
			fmt.Println("Yet to implement")
		default:
			fmt.Println("Invalid command")
		}
	}
}
