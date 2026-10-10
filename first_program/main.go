package main

import "fmt"

func main() {
	for {
		fmt.Println("\nВыбери программу:")
		fmt.Println("1. Угадай число")
		fmt.Println("0. Выход")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			guessNumber()
		case 0:
			fmt.Println("Пока!")
			return
		default:
			fmt.Println("Нет такого пункта")
		}
	}
}
