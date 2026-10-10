package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func passwordGenerator() {
	fmt.Println("генератор паролей")
	fmt.Println("какой длины пароль?")

	var line string
	fmt.Scanln(&line)

	length, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("введи число!")
		return
	}
	if length <= 0 {
		fmt.Println("длина должна быть больше 0")
		return
	}

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Println("добавить цифры? да/нет")
	var answer string
	fmt.Scanln(&answer)
	if answer == "да" {
		chars += "0123456789"
	}

	fmt.Println("добавить специальные символы? да/нет")
	fmt.Scanln(&answer)
	if answer == "да" {
		chars += "!@#$%^&*-_"
	}

	password := ""
	for i := 0; i < length; i++ {
		index := rand.IntN(len(chars))
		password += string(chars[index])
	}

	fmt.Println("твой пароль:", password)
}
