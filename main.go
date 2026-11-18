package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getAttempts() int {
	dif := 1
	fmt.Print("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n4. Custom attempts")
	fmt.Printf("\nEnter your choice: ")
	fmt.Scan(&dif)
	for dif > 4 || dif <= 0 {
		fmt.Println("Incorrect difficult. Choose between 1-4")
		fmt.Printf("\n\nEnter your choice: ")

		fmt.Scan(&dif)
	}
	switch {
	case dif == 1:
		return 10
	case dif == 2:
		return 5
	case dif == 3:
		return 3
	case dif == 4:
		var num int
		fmt.Printf("Enter the number of attempts: ")
		fmt.Scan(&num)
		return num
	}
	return dif
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.")
	attempts := getAttempts()

	count := 1
	var number int
	guess := rand.Intn(101)
	var guessed = false
	now := time.Now()
	for !guessed && count <= attempts {
		fmt.Printf("\nEnter your guess: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			panic(err)
		}
		if number == guess {
			guessed = true
			fmt.Println("Congratulations! You guessed the correct number in", count, "attempts. Time spent:", time.Since(now))
		}

		if number <= 0 || number > 100 {
			fmt.Println("Select a number between 1 and 100!")
		} else if number < guess && !guessed {
			fmt.Println("Incorrect! The number is greater than", number)
		} else if number > guess && !guessed {
			fmt.Println("Incorrect! The number is less than", number)
		}
		count++
	}

	if !guessed {
		fmt.Println("Runs out of chances. The chosen number is:", guess, "Time spent: ", time.Since(now))
	}

}
