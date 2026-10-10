package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("\nВыбери программу:")
		fmt.Println("1. Угадай число")
		fmt.Println("0. Выход")

		var line string
		fmt.Scanln(&line)

		choice, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("введи номер пункта цифрой")
			continue
		}

		switch choice {
		case 1:
			guessNumber()
		case 0:
			fmt.Println("пока!")
			return
		default:
			fmt.Println("нет такого пункта")
		}
	}
}
