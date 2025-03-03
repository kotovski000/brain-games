package main

import (
	"brain-games/games"
	"fmt"
)

func main() {
	fmt.Println("Choose a game:")
	fmt.Println("1. Find the smallest common multiple (LCM)")
	fmt.Println("2. Find the missing number in a geometric progression")
	fmt.Print("Your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		games.PlayLCM()
	case 2:
		games.PlayProgression()
	default:
		fmt.Println("Invalid choice!")
	}
}
