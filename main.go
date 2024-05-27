package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Seed the random number generator once at the beginning
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	numGuesses := 0

	for {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Remove newline from input
		input = input[:len(input)-1]

		guess, err := strconv.Atoi(input)
		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			continue
		}

		numGuesses++

		if guess < secretNumber {
			fmt.Println("Too low! Try a higher number.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try a lower number.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", numGuesses)
			break
		}
	}
}
