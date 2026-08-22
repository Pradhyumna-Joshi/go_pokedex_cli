package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()
		words := CleanInput(input)

		command := words[0]

		switch command {
		case "exit":
			if cmd, ok := Commands[command]; ok {
				if err := cmd.callback(); err != nil {
					fmt.Println(err)
					return
				}
			}
		case "help":
			if cmd, ok := Commands[command]; ok {
				if err := cmd.callback(); err != nil {
					fmt.Println(err)
					return
				}
				displayCommands()
			}
		default:
			fmt.Println("Unknown command")
		}

	}
}
