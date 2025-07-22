package archive

import (
	"fmt"
	"math/rand"
	"time"
)

// Структура для хранения параметров сложности
type Difficulty struct {
	Min      int    // Минимальное число диапазона
	Max      int    // Максимальное число диапазона
	Attempts int    // Количество попыток
	Name     string // Название уровня сложности
}

// Глобальные переменные
var (
	previousGuesses []int // Слайс для хранения истории попыток
	difficulties    = map[string]Difficulty{
		"easy":   {1, 50, 15, "Лёгкий"},
		"medium": {1, 100, 10, "Средний"},
		"hard":   {1, 200, 5, "Сложный"},
	}
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	// Выбор уровня сложности
	difficulty := selectDifficulty()

	// Генерация целевого числа
	target := rand.Intn(difficulty.Max-difficulty.Min+1) + difficulty.Min

	// Запуск игры
	runGame(target, difficulty)
}

// Функция выбора уровня сложности
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
			fmt.Println("Ошибка! Введите 1, 2 или 3")
		}
	}
}

// Основная функция игры
func runGame(target int, diff Difficulty) {
	previousGuesses = []int{} // Очистка истории при новой игре

	fmt.Printf("\n🎮 Игра началась! Угадай число от %d до %d. У тебя %d попыток.\n",
		diff.Min, diff.Max, diff.Attempts)

	// Основной игровой цикл
	for attempt := 1; attempt <= diff.Attempts; attempt++ {
		// Получаем число от игрока
		guess := getPlayerInput(attempt, diff.Attempts)

		// Проверяем угадал ли игрок
		if isCorrectGuess(guess, target) {
			return // Выход если угадал
		}
	}

	// Если попытки закончились
	fmt.Printf("\n😢 Попытки закончились. Загаданное число: %d\n", target)
	showHistory()
}

// Функция для ввода числа игроком
func getPlayerInput(attempt, maxAttempts int) int {
	for {
		fmt.Printf("\n👉 Попытка %d/%d: Введите число: ", attempt, maxAttempts)
		var num int

		// Проверка корректности ввода
		if _, err := fmt.Scan(&num); err == nil {
			// Сохраняем попытку в историю
			previousGuesses = append(previousGuesses, num)
			return num
		}

		// Обработка ошибки ввода
		fmt.Println("Ошибка! Введите целое число.")
		var discard string
		fmt.Scanln(&discard) // Очистка буфера ввода
	}
}

// Проверка правильности числа
func isCorrectGuess(guess, target int) bool {
	// Определяем направление подсказки
	direction := ""
	if guess < target {
		direction = "больше"
	} else if guess > target {
		direction = "меньше"
	} else {
		// Игрок угадал число
		fmt.Println("\n🎉 Поздравляю! Ты угадал число!")
		showHistory()
		return true
	}

	// Получаем подсказку по расстоянию
	hint := getDistanceHint(guess, target)
	fmt.Printf("%s Загаданное число %s.", hint, direction)
	showHistory()

	return false
}

// Показать историю попыток
func showHistory() {
	if len(previousGuesses) == 0 {
		return
	}

	fmt.Print("\n📋 История попыток: ")
	for _, g := range previousGuesses {
		fmt.Printf("%d ", g)
	}
	fmt.Println()
}

// Генерация подсказки по расстоянию до числа
func getDistanceHint(guess, target int) string {
	distance := abs(guess - target)

	switch {
	case distance == 0:
		return "🎉"
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

// Вычисление модуля числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
ПОЯСНЕНИЯ ПО СТРУКТУРЕ:

=== НАСТРОЙКА ИГРЫ ===
Difficulty - структура, хранящая параметры уровня сложности:
  - Min: минимальное число диапазона
  - Max: максимальное число диапазона
  - Attempts: количество попыток
  - Name: название уровня сложности

selectDifficulty() - функция выбора уровня сложности:
  - Отображает меню выбора
  - Возвращает выбранную сложность

=== ИГРОВОЙ ПРОЦЕСС ===
runGame() - основной игровой цикл:
  - Управляет процессом угадывания
  - Обрабатывает окончание игры

getPlayerInput() - обработка ввода пользователя:
  - Запрашивает число у игрока
  - Проверяет корректность ввода
  - Сохраняет попытку в историю

isCorrectGuess() - проверка числа игрока:
  - Сравнивает с загаданным числом
  - Дает подсказки (больше/меньше)
  - Определяет победу

=== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ===
showHistory() - отображает историю попыток:
  - Показывает все предыдущие числа

getDistanceHint() - генерирует подсказки:
  - "Горячо/Холодно" на основе расстояния

abs() - вычисляет модуль числа:
  - Используется для определения расстояния

ОСОБЕННОСТИ ПРОГРАММЫ:
- Полная история всех попыток
- Интеллектуальные подсказки по расстоянию
- Защита от неверного ввода
- Чистый и понятный код
- Гибкая система уровней сложности

ИНСТРУКЦИЯ ПО ИСПОЛЬЗОВАНИЮ:
1. Сохраните код в файл guess_number.go
2. Запустите командой: go run guess_number.go
3. Выберите уровень сложности (1-3)
4. Используйте подсказки для угадывания числа
5. Для победы угадайте число за отведенное количество попыток
*/
