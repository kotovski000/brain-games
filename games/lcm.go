package games

import (
	"fmt"
	"math/rand"
	"time"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmOfThree(a, b, c int) int {
	return lcm(lcm(a, b), c)
}

func PlayLCM() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("Find the smallest common multiple of given numbers.")

	correctAnswersNeeded := 3
	correctAnswers := 0

	for correctAnswers < correctAnswersNeeded {
		num1 := rand.Intn(20) + 1
		num2 := rand.Intn(20) + 1
		num3 := rand.Intn(20) + 1
		correctAnswer := lcmOfThree(num1, num2, num3)

		fmt.Printf("Question: %d %d %d\n", num1, num2, num3)
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
