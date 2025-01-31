# Number Guessing Game

A simple number guessing game written in Go. The player must guess a randomly chosen number within a limited number of attempts based on the selected difficulty level.

https://roadmap.sh/projects/number-guessing-game

## Features
- Three difficulty levels: Easy, Medium, and Hard.
- Random number generation between 1 and 100.
- Input validation to ensure only numbers are entered.
- Feedback on whether the guess was too high or too low.
- Exit option to quit the game at any time.

## Installation

Ensure you have Go installed on your system.

Clone this repository or create a new directory and add the `main.go` file:

```sh
git clone https://github.com/trinverdale-bit/number-guessing-game.git
cd number-guessing-game
```

## Usage

Run the program using:

```sh
go run main.go
```

You will be prompted to select a difficulty level:

```sh
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances) - Default
3. Hard (3 chances)
Enter your choice (or type 'exit' to quit):
```

After selecting a difficulty, you will be asked to guess a number.

Example gameplay:

```
Enter your guess: (10 attempts remaining): 50
Incorrect! The number is more than 50
Enter your guess: (9 attempts remaining): 75
Incorrect! The number is less than 75
...
Congratulations! You guessed the correct number in 4 attempts.
```

If you fail to guess within the attempts limit:

```
Uh oh! You ran out of attempts! The correct number was: 42
```

## Code Overview

### `gameLoop()`
Handles the main menu and difficulty selection.

### `easyDifficulty()`, `mediumDifficulty()`, `hardDifficulty()`
Each function contains logic for the respective difficulty, generating a random number and checking user guesses.
