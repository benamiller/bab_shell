package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var _ = fmt.Fprint

func main() {
	reader := bufio.NewReader(os.Stdin)

	builtinTypes := map[string]string{
		"exit": "exit is a shell builtin",
		"echo": "echo is a shell builtin",
		"type": "type is a shell builtin",
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		input = strings.TrimSpace(input)

		commands := strings.Fields(input)

		switch commands[0] {
		case "exit":
			code, err := strconv.Atoi(commands[1])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))
		case "type":
			type_queried, exists := builtinTypes[commands[1]]
			if exists {
				fmt.Println(type_queried)
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", commands[1])
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		}
	}
}
