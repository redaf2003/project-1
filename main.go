package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

//var previousGuesses []int

func main() {
	rand.Seed(time.Now().UnixNano())
	diff := selectDifficulty()
	targetNumber := rand.Intn(diff.Max-diff.Min+1) + diff.Min
	//(diff.Max-diff.Min+1) + diff.Min

	fmt.Printf("\n🎮 Игра началась! Угадай число от %d до %d. У тебя %d попыток.\n",
		diff.Min, diff.Max, diff.Attempts)
	//fmt.Printf("\nИгра началась! Угадай число от %d до %d. У тебя %d попыток.\n",
	//	diff.Min, diff.Max, diff.Attempts)

	RunGame(targetNumber, diff.Attempts)
	//targetNumber := Numbergeneration(1, 100)
	//maxAttepts := 10
	//RunGame(targetNumber, maxAttepts)
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

func RunGame(targetNumber, maxAttempts int) {
	previousGuesses = []int{}

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		guess := getPlayerInput(attempt, maxAttempts)

		if checkGuess(guess, targetNumber) {
			return
		}
	}

	fmt.Printf("\n😢 Попытки закончились. Загаданное число: %d\n", targetNumber)
	showPreviousGuesses()
}

//fmt.Printf("Игра началась ❗ Угадай число от 1 до 100. У тебя будет %d попыток.\n", maxAttepts)

//for attemmpt := 1; attemmpt <= maxAttepts; attemmpt++ {

//	gues := getPlayerGuess(attemmpt)
//	clue := getDistanceHint(gues, targetNumber)

//	if gues == targetNumber {
//		fmt.Printf("🎉 Поздравляю! Ты угадал число!")
//		showPreviousGuesses()
//		break
//	}

//	if gues < targetNumber {
//		fmt.Printf("%s Загаданное число больше.\n", clue)
//	} else {
//		fmt.Printf("%s Загаданное число меньше.\n", clue)
//	}

// }
// fmt.Printf("Попытки закончились. Загаданное число было: %d\n", targetNumber)
// }
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

//func getPlayerGuess(attempt int) int {
//	for {
//		fmt.Printf("Попытка %d. Ведите число:", attempt)
//		guess := 0
//		if _, err := fmt.Scan(&guess); err == nil {
//			previousGuesses = append(previousGuesses, guess)
//			return guess
//		}
//		fmt.Println("Ошибка! Пожалуйста, введите целое число.")
//		var discard string
//		fmt.Scanln(&discard)
//	}
//}

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
