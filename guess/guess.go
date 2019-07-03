package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please think of a number between 1 to 100.")
	fmt.Println("Press ENTER when ready.")
	scanner.Scan()

	low := 1
	high := 100

	for {
		guess := (low + high) / 2
		fmt.Println("I guess the number is:", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) Too High?")
		fmt.Println("(b) Too Low?")
		fmt.Println("(c) Correct?")
		scanner.Scan()

		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won")
			break
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
