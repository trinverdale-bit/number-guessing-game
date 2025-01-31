package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func gameLoop() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	for {
		fmt.Println("\nPlease select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances) - Default")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice (or type 'exit' to quit): ")

		var difficulty string
		fmt.Scanln(&difficulty)
		difficulty = strings.ToLower(strings.TrimSpace(difficulty))

		switch difficulty {
		case "1":
			easyDifficulty()
		case "2":
			mediumDifficulty()
		case "3":
			hardDifficulty()
		case "":
			mediumDifficulty()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input. Please select a valid number (1, 2, 3) or type 'exit' to quit.")
		}
	}
}

func easyDifficulty() {
	fmt.Println("Great! You have selected the Easy difficulty level.")
	fmt.Println("Let's start the game!")
    fmt.Println()

	randomNumber := rand.Intn(101)
	attempts := 0
	attempsCap := 10

	for attempts < attempsCap {
        fmt.Printf("Enter your guess: (%d attempts remaining): ", attempsCap-attempts)

		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
            continue
		}

        attempts++

		switch {
		case guess < randomNumber:
			fmt.Printf("Incorrect! The number is more than %d\n", guess)
		case guess > randomNumber:
			fmt.Printf("Incorrect! The number is less than %d\n", guess)
		case guess == randomNumber:
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}
	}

    fmt.Println("Uh oh! You ran out of attempts! The correct number was:", randomNumber)
}

func mediumDifficulty() {
    fmt.Println("Great! You have selected the Medium difficulty level.")
	fmt.Println("Let's start the game!")
    fmt.Println()

	randomNumber := rand.Intn(101)
	attempts := 0
	attempsCap := 5

	for attempts < attempsCap {
        fmt.Printf("Enter your guess: (%d attempts remaining): ", attempsCap-attempts)

		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
            continue
		}

        attempts++

		switch {
		case guess < randomNumber:
			fmt.Printf("Incorrect! The number is more than %d\n", guess)
		case guess > randomNumber:
			fmt.Printf("Incorrect! The number is less than %d\n", guess)
		case guess == randomNumber:
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}
	}

    fmt.Println("Uh oh! You ran out of attempts! The correct number was:", randomNumber)
}

func hardDifficulty() {
    fmt.Println("Great! You have selected the Easy difficulty level.")
	fmt.Println("Let's start the game!")
    fmt.Println()

	randomNumber := rand.Intn(101)
	attempts := 0
	attempsCap := 10

	for attempts < attempsCap {
        fmt.Printf("Enter your guess: (%d attempts remaining): ", attempsCap-attempts)

		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
            continue
		}

        attempts++

		switch {
		case guess < randomNumber:
			fmt.Printf("Incorrect! The number is more than %d\n", guess)
		case guess > randomNumber:
			fmt.Printf("Incorrect! The number is less than %d\n", guess)
		case guess == randomNumber:
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}
	}

    fmt.Println("Uh oh! You ran out of attempts! The correct number was:", randomNumber)
}

func main() {
	gameLoop()
}
