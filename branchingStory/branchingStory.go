package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.noPath != nil && node.yesPath != nil {

		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			answer := scanner.Text()
			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("That answer was not an option! Please answer in yes or no!!")
			}
		}
	}
}

func main() {
	root := storyNode{"You are at the entrance to a dark cave. Do you want to go in?", nil, nil}
	winning := storyNode{"You have won!", nil, nil}
	losing := storyNode{"You lost!", nil, nil}
	root.yesPath = &losing
	root.noPath = &winning
	root.play()
}
