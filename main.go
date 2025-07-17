package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	guess, attempt := 0, 10

	for i := 1; i <= attempt; i++ {
		fmt.Printf(" Игра началась! Желаю успехов:Gопытка %d: ведите ваше число ", i)
		fmt.Scan(&guess)
		if guess < randomNumber {
			fmt.Println("Парень ты вел маленькое число ")
		} else if guess > randomNumber {
			fmt.Println("Парень ты дурак ты вел большое число ты как Егор дебил ")
		} else {
			fmt.Println("Красава ты угадал с числом можешь даже покрутить казик если хочешь ")
			break
		}
		if i == 0 {
			fmt.Printf("Ты исчерпал все попытки, ты лошара как Егор число было: %d\n", randomNumber)
		}

	}

}
