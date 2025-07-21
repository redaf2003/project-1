package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	targetNumber := Numbergeneration(1, 100)
	maxAttepts := 10
	RunGame(targetNumber, maxAttepts)
}
func Numbergeneration(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func RunGame(targetNumber, maxAttepts int) {
	fmt.Printf("Игра началась ❗ Угадай число от 1 до 100. У тебя будет %d попыток.\n", maxAttepts)
	for attemmpt := 1; attemmpt <= maxAttepts; attemmpt++ {
		gues := getPlayerGuess(attemmpt)
		switch {
		case gues < targetNumber:
			fmt.Println("Твое число слишком маленькое! ")
		case gues > targetNumber:
			fmt.Println("Твое число слишком большое ")
		default:
			fmt.Println("Красава ты победил ")
			return
		}
	}
	fmt.Printf("Попытки закончились. Загаданное число было: %d\n", targetNumber)
}

func getPlayerGuess(attempt int) int {
	for {
		fmt.Printf("Попытка %d. Ведите число:", attempt)
		guess := 0
		if _, err := fmt.Scan(&guess); err == nil {
			return guess
		}
		fmt.Println("Ошибка! Пожалуйста, введите целое число.")
		var discard string
		fmt.Scanln(&discard)
	}
}

//rand.Seed(time.Now().UnixNano())
//randomNumber := rand.Intn(100) + 1
//guess, attempt := 0, 10

//for i := 1; i <= attempt; i++ {
//	fmt.Printf(" Игра началась! Желаю успехов:Попытка %d: ведите ваше число ", i)
//	_, err := fmt.Scan(&guess)
//	if err != nil {

//		fmt.Println("Ошибка при воде : Ведите конретное число !!!")
//		fmt.Scan()
//		i--
//		continue

//	}

//		if guess < randomNumber {
//			fmt.Println("Парень ты вел маленькое число ")
//		} else if guess > randomNumber {
//			fmt.Println("Парень ты вел большое число ")
//		} else {
//			fmt.Println("Красава ты угадал с числом можешь даже покрутить казик если хочешь ")
//			break
//		}
//		if i == 0 {
//			fmt.Printf("Ты исчерпал все попытки, число  было: %d\n", randomNumber)
//		}
//
//	}

//}
