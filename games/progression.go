package games

import (
	"fmt"
	"math/rand"
	"time"
)

func generateProgression(start, step, length int) []int {
	progression := make([]int, length)
	for i := 0; i < length; i++ {
		progression[i] = start * intPow(step, i)
	}
	return progression
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func PlayProgression() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("What number is missing in the progression?")

	correctAnswersNeeded := 3
	correctAnswers := 0

	for correctAnswers < correctAnswersNeeded {
		start := rand.Intn(10) + 1
		step := rand.Intn(4) + 2
		length := rand.Intn(6) + 5
		progression := generateProgression(start, step, length)
		hiddenIndex := rand.Intn(length)
		correctAnswer := progression[hiddenIndex]
		progression[hiddenIndex] = -1

		fmt.Print("Question: ")
		for _, num := range progression {
			if num == -1 {
				fmt.Print(".. ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()

		fmt.Print("Your answer: ")
		var userAnswer int
		fmt.Scanln(&userAnswer)

		if userAnswer == correctAnswer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", userAnswer, correctAnswer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}
	}

	fmt.Printf("Congratulations, %s!\n", name)
}
