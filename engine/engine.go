package engine

import (
	"fmt"
)

type GameLogic interface {
	GenerateQuestion() (string, int)
	GetInstructions() string
}

func RunGame(game GameLogic) {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")
	var name string
	if _, err := fmt.Scanln(&name); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println(game.GetInstructions())

	correctAnswersNeeded := 3
	correctAnswers := 0

	for correctAnswers < correctAnswersNeeded {
		question, correctAnswer := game.GenerateQuestion()
		fmt.Printf("Question: %s\n", question)
		fmt.Print("Your answer: ")
		var userAnswer int
		if _, err := fmt.Scanln(&userAnswer); err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		if userAnswer == correctAnswer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("'%d' is wrong answer :(. Correct answer was '%d'.\n", userAnswer, correctAnswer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}
	}

	fmt.Printf("Congratulations, %s!\n", name)
}
