package main

import (
	"fmt"
	"math/rand"
)

var numberToGuess int = rand.Intn(100)
var difficultyMap map[int]int = make(map[int]int)
var difficultyStringMap map[int]string = make(map[int]string)

func init() {
	difficultyMap[1] = 10
	difficultyMap[2] = 5
	difficultyMap[3] = 3

	difficultyStringMap[1] = "Easy"
	difficultyStringMap[2] = "Medium"
	difficultyStringMap[3] = "Hard"

}

func main() {

	// argsWithFile := os.Args

	// argsWithoutFilename := argsWithFile[1:]
	var option int
	startGame()
	fmt.Print("Enter your choice:")
	difficultyChoice, err := fmt.Scanln(&option)
	if err != nil || difficultyChoice > 1 {
		fmt.Println("No choice or invalid number of choices given! Exiting")
		return
	}

	fmt.Printf("Great! You've chosed %s difficulty\n\n", difficultyStringMap[option])
	gameLoop(difficultyMap[option])

}

func startGame() {
	fmt.Println("Welcome to Number Guessing game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println()
	fmt.Println("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
}

func gameLoop(n int) {

	var userEnteredValue int
	for ind := range n {
		fmt.Print("Enter your guess:")
		_, err := fmt.Scanln(&userEnteredValue)
		if err != nil {
			fmt.Println("Invalid Value Entered!")
		}
		_ = ind
		if userEnteredValue == numberToGuess {
			fmt.Println("Congrats! You've guessed the number")
			fmt.Println()
			return
		} else if userEnteredValue < numberToGuess {
			fmt.Println("Incorrect! The number is more than ", userEnteredValue)
			fmt.Println()
		} else if userEnteredValue > numberToGuess {
			fmt.Println("Incorrect! The number is less than ", userEnteredValue)
			fmt.Println()
		}
	}

	fmt.Println("Oop! Couldn't get it this time")
	fmt.Println()

}
