package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	m "github.com/iamvishnuavenu/carton/pkg/meta"
	st "github.com/iamvishnuavenu/carton/pkg/statement"
)

const (
	PREPARE_SUCCESS               = 0
	PREPARE_UNRECONIZED_STATEMENT = 1
)

func print_promt() {
	fmt.Print("carton > ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		print_promt()
		command, _ := reader.ReadString('\n')
		// stripping new line
		command = strings.TrimSuffix(command, "\n")

		// Handling Meta command
		if command[0] == '.' {
			switch m.DoMetaCommamd(command) {
			case m.META_COMMAND_SUCCESS:
				continue
			case m.META_COMMAND_UNRECOGNIZED:
				fmt.Printf(" Unrecognized command: %s \n", command)
				continue
			}
		}

		var statement st.Statement
		switch st.PrepareStatement(command, &statement) {
		case PREPARE_SUCCESS:
		case PREPARE_UNRECONIZED_STATEMENT:
			fmt.Printf(" Unrecognized command: %s \n", command)
			continue
		}

		st.ExecuteStatement(&statement)
		fmt.Println("Executed.")
	}
}
