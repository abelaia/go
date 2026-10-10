package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func guessNumber() {
	fmt.Println("игра «Угадай число»")
	fmt.Println("я загадал число от 1 до 100, попробуй угадать:")

	attempts := 0
	randomNum := rand.IntN(100) + 1
	for {
		var line string
		fmt.Scanln(&line)

		input, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("вводи только числа!")
			continue
		}

		attempts++

		if randomNum == input {
			fmt.Printf("угадано! попыток: %d", attempts)
			break
		} else if randomNum < input {
			fmt.Println("меньше")
		} else {
			fmt.Println("больше")
		}
	}
}
