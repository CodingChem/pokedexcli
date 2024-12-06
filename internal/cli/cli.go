package cli

import (
	"bufio"
	"fmt"
	"os"
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
		command := scanner.Text()
		CommandMap[command].callback()
	}
}
