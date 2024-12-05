package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Fprint

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		command = strings.TrimSpace(command)

		if command == "exit 0" {
			break
		}

		fmt.Println(command + ": not found")
	}
}
