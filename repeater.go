package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type some text to be repeated.")
	fmt.Print("Text: ")

	text, err := reader.ReadString('\n')
	if nil != err {
		log.Fatalf("Error while reading text entry: %s", err)
	}
	line := strings.TrimSpace(text)

	fmt.Print("Number of times to repeat: ")

	repeats, err := reader.ReadString('\n')
	if nil != err {
		log.Fatalf("Error while reading repeat count: %s", err)
	}

	count, err := strconv.Atoi(strings.TrimSpace(repeats))
	if nil != err {
		log.Fatalf("Error while converting int: %s", err)
	}

	repeat(line, count)
}

func repeat(line string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(line)
	}
}
