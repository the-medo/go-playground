package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0

	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		}

		text := strings.TrimSpace(scanner.Text())

		switch text {
		case CmdHello:
			numCommands++
			fmt.Println("Command response: Hi!")
		case CmdGoodbye:
			numCommands++
			fmt.Println("Command response: Bye!")
		}
		if text != "" {
			numLines++
		}
	}

	fmt.Printf("You entered %v lines \n", numLines)
	fmt.Printf("You entered %v commands \n", numCommands)
}
