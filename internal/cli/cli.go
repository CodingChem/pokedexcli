package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex->")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println("encountered fatal error, shutting down...")
			return
		}
		command, arg, err := parseInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		if cmd, exists := CommandMap[command]; exists {
			if err := cmd.callback(arg); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Invalid command:", command)
		}
	}
}

func parseInput(input string) (command string, arg string, err error) {
	splitInput := strings.Split(input, " ")
	argLen := len(splitInput)
	switch {
	case argLen == 0:
		return "", "", fmt.Errorf("Error: no input to parse?")
	case argLen == 1:
		command = splitInput[0]
		arg = ""
		err = nil
	case argLen == 2:
		command = splitInput[0]
		arg = strings.Join(splitInput[1:], "")
		err = nil
	}
	return
}
