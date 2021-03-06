package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, ch := range node.choices {
			fmt.Println(ch.cmd, ":", ch.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, ch := range node.choices {
		if strings.ToLower(ch.cmd) == strings.ToLower(cmd) {
			return ch.nextNode
		}
	}
	fmt.Println("Sorry I didn't understand that!!")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in a large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south a passage appears to head upward. The eastern passages appears
	flat and well travelled.`}

	darkRoom := storyNode{text: "It is pitch black and you cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark room is now lit by your lantern. You can continue north or head back south"}

	grue := storyNode{text: "While stumbling around in darkness, you are eaten by a grue."}

	trap := storyNode{text: "You head down the well travelled path when suddenly a trap door opens and you fall into a pit."}

	treasure := storyNode{text: "You arrived at a small chamber, filled with treasure."}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on the lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End!")

}
