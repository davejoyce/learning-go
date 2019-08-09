package main

import (
	"bufio"
	"fmt"
	"github.com/davejoyce/learning-go/pkg/guess"
	"github.com/davejoyce/learning-go/pkg/repeater"
	"log"
	"os"
	"strings"
)

func main() {
	var choice string = "guess"

	fmt.Println("Pick a thing to do:")
	fmt.Println("[guess] a number")
	fmt.Println("[repeat] something")
	fmt.Print("choice [", choice, "]: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if nil != err {
		log.Fatalf("Failed to read choice: %s", err)
	}
	choice = strings.TrimSpace(input)
	fmt.Println("You chose", choice)

	switch choice {
	case "repeat":
		repeater.Repeater()
	case "guess":
		guess.Guess()
	default:
		guess.Guess()
	}

	fmt.Println("----------------")
	fmt.Println("DONE")
}
