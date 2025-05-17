package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		secret := rand.Intn(100) + 1
		fmt.Println("\nGuess the number (between 1 and 100):")

		for {
			guess := getValidIntegerInput("\nEnter your guess: ")

			if guess < secret {
				fmt.Println("Too low! Try again.")
			} else if guess > secret {
				fmt.Println("Too high! Try again.")
			} else {
				fmt.Println("\n🎉 Correct! You guessed it!")
				break
			}
		}

		fmt.Print("\nDo you want to guess another number? (yes/no): ")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "yes" {
			fmt.Println("\nThanks for playing!")
			break
		}
	}
}

func getValidIntegerInput(prompt string) int {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		// Reject decimal inputs
		if strings.Contains(input, ".") {
			fmt.Println("\nOnly whole numbers are allowed.")
			continue
		}

		// Try converting to integer
		value, err := strconv.Atoi(input)
		if err != nil || value < 1 || value > 100 {
			fmt.Println("\nInvalid input. Please enter a number between 1 and 100.")
			continue
		}

		return value
	}
}
