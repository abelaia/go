package main

import (
	"fmt"
	"strconv"
)

func converter() {
	for {
		fmt.Println("конвертер")
		fmt.Println("1. Цельсий → Фаренгейт")
		fmt.Println("2. Фаренгейт → Цельсий")
		fmt.Println("3. Километры → Мили")
		fmt.Println("0. Назад")

		var line string
		fmt.Scanln(&line)

		choice, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("введите номер пункта цифрой!")
			continue
		}

		switch choice {
		case 1:
			celsiusToFahrenheit()
		case 2:
			fahrenheitToCelsius()
		case 3:
			kilometersToMiles()
		case 0:
			return
		default:
			fmt.Println("нет такого пункта")
		}
	}
}

func celsiusToFahrenheit() {
	var F float64
	var line string
	fmt.Println("введи температуру в °C: ")
	fmt.Scanln(&line)

	C, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println("введите число!")
		return
	}

	F = C*9/5 + 32
	fmt.Printf("в %.2f градусах Цельсия %.2f градусов по Фаренгейту\n", C, F)
}

func fahrenheitToCelsius() {
	var C float64
	var line string
	fmt.Println("введи температуру в °F: ")
	fmt.Scanln(&line)

	F, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println("введите число!")
		return
	}

	C = (F - 32) * 5 / 9
	fmt.Printf("в %.2f градусах по Фаренгейту %.2f градусов Цельсия\n", F, C)
}

func kilometersToMiles() {
	var kilometers float64
	var line string
	fmt.Println("введи количество километров: ")
	fmt.Scanln(&line)

	kilometers, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println("введите число!")
		return
	}

	miles := kilometers * 0.621371
	fmt.Printf("%.2f километров в %.2f милях\n", kilometers, miles)
}
