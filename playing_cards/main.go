/**
 * Author: SAHIL SHARMA
 * Created On: March 01, 2018
 * Language: GO
 * Project: 
 */


package main

import "fmt"

func main() {
	// Two ways to define a NEW variables. := won't work when reintializing the variable again.

	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Spades"

	fmt.Println(card)
}