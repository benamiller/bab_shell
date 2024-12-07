package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var _ = fmt.Fprint

func find_directory(command string) (string, bool) {
	path := os.Getenv("PATH")
	directories := strings.Split(path, ":")

	for _, directory := range directories {
		directory = strings.TrimSuffix(directory, "/")
		candidatePath := filepath.Join(directory, command)

		info, err := os.Stat(candidatePath)
		if err == nil && !info.IsDir() {
			return candidatePath, true
		}
	}
	return command, false
}

func run_command(command string, args []string) (string, bool) {
	path, exists := find_directory(command)
	if !exists {
		return "Command not found", false
	}

	cmd := exec.Command(path, args...)

	output, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(output)), true
	}

	return cmd.Path, true
}

func change_directory(path string) error {
	if path == "~" {
		path = os.Getenv("HOME")
	}

	err := os.Chdir(path)

	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
		return err
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	builtinTypes := map[string]string{
		"exit": "exit is a shell builtin",
		"echo": "echo is a shell builtin",
		"type": "type is a shell builtin",
		"pwd":  "pwd is a shell builtin",
		"cd":   "cd is a shell builtin",
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
				command := commands[1]
				pathToCommand, exists := find_directory(command)
				if exists {
					fmt.Fprintf(os.Stdout, "%s is %s\n", command, pathToCommand)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", command)
				}
			}
		case "pwd":
			directory, err := os.Getwd()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println(directory)
		case "cd":
			change_directory(commands[1])
		default:
			path, exists := run_command(commands[0], commands[1:])
			if !exists {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
			} else {
				fmt.Println(path)
			}
		}
	}
}
