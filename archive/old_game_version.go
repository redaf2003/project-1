package archive

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target := generateRandomNumber(1, 100)
	attempts := 10
	runGame(target, attempts)
}

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func runGame(target, maxAttempts int) {
	fmt.Printf("🎮 Игра началась! Угадай число от 1 до 100. Попыток: %d\n", maxAttempts)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		guess := getPlayerInput(attempt)

		if checkGuess(guess, target) {
			return
		}
	}

	fmt.Printf("😢 Попытки закончились. Загаданное число: %d\n", target)
}

func getPlayerInput(attempt int) int {
	for {
		fmt.Printf("👉 Попытка %d: Введите число: ", attempt)
		var num int
		if _, err := fmt.Scan(&num); err == nil {
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
		fmt.Println("Твое число слишком маленькое!")
		return false
	case guess > target:
		fmt.Println("Твое число слишком большое!")
		return false
	default:
		fmt.Println("🎉 Красава, ты победил!")
		return true
	}
}
