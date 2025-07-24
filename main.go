package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameResult struct {
	Date        string `json:"date"`
	Outcome     string `json:"outcome"`
	Attempts    int    `json:"attempts"`
	Target      int    `json:"target"`
	Difficultys string `json:"difficulty"`
}

type Difficulty struct {
	Min      int
	Max      int
	Attempts int
	Name     string
}

var (
	previousGuesses []int
	difficulties    = map[string]Difficulty{
		"easy":   {1, 50, 15, "Лёгкий"},
		"medium": {1, 100, 10, "Средний"},
		"hard":   {1, 200, 5, "Сложный"},
	}
)

func main() {

	rand.Seed(time.Now().UnixNano())

	for {
		diff := selectDifficulty()
		targetNumber := rand.Intn(diff.Max-diff.Min+1) + diff.Min

		fmt.Printf("\n🎮 Игра началась! Угадай число от %d до %d. У тебя %d попыток.\n",
			diff.Min, diff.Max, diff.Attempts)

		RunGame(targetNumber, diff.Attempts, diff.Name)

		if askForReplay() == false {
			fmt.Println("\nСпасибо за игру! До свидания!")
			break
		}
	}

}

func saveResult(won bool, attempts, target int, difficulty string) {
	result := GameResult{
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		Outcome:     "lose",
		Attempts:    attempts,
		Target:      target,
		Difficultys: difficulty,
	}
	if won {
		result.Outcome = "win"
	}

	var results []GameResult
	file, err := os.ReadFile("results.json")
	if err == nil {
		json.Unmarshal(file, &results)
	}

	results = append(results, result)

	data, _ := json.MarshalIndent(results, "", "  ")
	os.WriteFile("results.json", data, 0644)
}

func askForReplay() bool {
	for {
		fmt.Print("\nХотите сыграть ещё раз? (да/нет): ")
		var answer string
		fmt.Scan(&answer)

		switch answer {
		case "да", "д", "yes", "y":
			return true
		case "нет", "н", "no", "n":
			return false
		default:
			fmt.Println("Пожалуйста, введите 'да' или 'нет'")
		}
	}
}

func selectDifficulty() Difficulty {
	for {
		fmt.Print(`
Выберите сложность:
1. Лёгкий (1-50, 15 попыток)
2. Средний (1-100, 10 попыток)
3. Сложный (1-200, 5 попыток)
Ваш выбор: `)

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1", "easy":
			return difficulties["easy"]
		case "2", "medium":
			return difficulties["medium"]
		case "3", "hard":
			return difficulties["hard"]
		default:
			fmt.Println("⚠️ Неверный ввод. Попробуйте снова.")
		}
	}
}

func Numbergeneration(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RunGame(targetNumber, maxAttempts int, difDifficulty string) bool {
	previousGuesses = []int{}
	usedAttempts := 0
	won := false

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		guess := getPlayerInput(attempt, maxAttempts)

		if checkGuess(guess, targetNumber) {
			won = true
			break
		}
	}

	if !won {
		fmt.Printf("\n😢 Попытки закончились. Загаданное число: %d\n", targetNumber)
	}
	showPreviousGuesses()

	saveResult(won, usedAttempts, targetNumber, difDifficulty)
	return won
}

func getPlayerInput(attempt, maxAttempts int) int {
	for {
		fmt.Printf("\n👉 Попытка %d/%d: Введите число: ", attempt, maxAttempts)
		var num int
		if _, err := fmt.Scan(&num); err == nil {
			previousGuesses = append(previousGuesses, num)
			return num
		}

		fmt.Println("⚠️ Ошибка! Введите целое число.")
		var discard string
		fmt.Scanln(&discard)
	}
}

func checkGuess(guess, target int) bool {
	switch {
	case guess < target:
		fmt.Printf("%s Загаданное число больше.", getDistanceHint(guess, target))
	case guess > target:
		fmt.Printf("%s Загаданное число меньше.", getDistanceHint(guess, target))
	default:
		fmt.Println("\n🎉 Поздравляю! Ты угадал число!")
		showPreviousGuesses()
		return true
	}

	showPreviousGuesses()
	return false
}

func showPreviousGuesses() {
	if len(previousGuesses) == 0 {
		return
	}

	fmt.Print("\n📋 Ваши попытки: ")
	for _, g := range previousGuesses {
		fmt.Printf("%d ", g)
	}
	fmt.Println()
}

func getDistanceHint(gues, targetNumber int) string {
	distance := abs(gues - targetNumber)
	switch {
	case distance == 0:
		return "🎉 "
	case distance <= 3:
		return "🔥 ОЧЕНЬ ГОРЯЧО!"
	case distance <= 7:
		return "🔥 Горячо!"
	case distance <= 15:
		return "🙂 Тепло"
	default:
		return "❄️ Холодно"
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
