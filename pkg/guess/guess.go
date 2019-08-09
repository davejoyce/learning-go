package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const MAX_GUESSES = 10

func Guess() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < MAX_GUESSES; guesses++ {
		fmt.Println("You have", MAX_GUESSES-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if nil != err {
			log.Fatal(err)
		}
		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if nil != err {
			log.Fatal(err)
		}

		fmt.Println("You guessed", guess)
		if guess < target {
			fmt.Println("Oops... your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops... your guess was HIGH")
		} else {
			fmt.Println("You nailed it!! It was", target)
			break
		}
	}
}
