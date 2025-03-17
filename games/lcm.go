package games

import (
	"brain-games/engine"
	"brain-games/utils"
	"fmt"
)

type LCMGame struct{}

func (g LCMGame) GenerateQuestion() (string, int) {
	num1 := utils.RandomInt(20) + 1
	num2 := utils.RandomInt(20) + 1
	num3 := utils.RandomInt(20) + 1
	correctAnswer := lcmOfThree(num1, num2, num3)
	question := fmt.Sprintf("%d %d %d", num1, num2, num3)
	return question, correctAnswer
}

func (g LCMGame) GetInstructions() string {
	return "Find the smallest common multiple of given numbers."
}

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
	game := LCMGame{}
	engine.RunGame(game)
}
