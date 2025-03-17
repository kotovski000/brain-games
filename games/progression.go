package games

import (
	"brain-games/engine"
	"brain-games/utils"
	"fmt"
)

type ProgressionGame struct{}

func (g ProgressionGame) GenerateQuestion() (string, int) {
	start := utils.RandomInt(10) + 1
	step := utils.RandomInt(4) + 2
	length := utils.RandomInt(6) + 5

	progression := generateProgression(start, step, length)
	hiddenIndex := utils.RandomInt(int64(length))
	correctAnswer := progression[hiddenIndex]
	progression[hiddenIndex] = -1

	question := ""
	for _, num := range progression {
		if num == -1 {
			question += ".. "
		} else {
			question += fmt.Sprintf("%d ", num)
		}
	}

	return question, correctAnswer
}

func (g ProgressionGame) GetInstructions() string {
	return "What number is missing in the progression?"
}

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
	game := ProgressionGame{}
	engine.RunGame(game)
}
