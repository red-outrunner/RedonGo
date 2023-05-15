package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 1 and 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")

	var guess int
	numGuesses := 0

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
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
