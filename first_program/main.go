package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("\nвыбери программу:")
		fmt.Println("1. угадай число")
		fmt.Println("2. конвертер")
		fmt.Println("3. генератор паролей")
		fmt.Println("0. выход")

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
		case 2:
			converter()
		case 3:
			passwordGenerator()
		case 0:
			fmt.Println("пока!")
			return
		default:
			fmt.Println("нет такого пункта")
		}
	}
}
