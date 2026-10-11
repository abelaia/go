package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("список пуст")
		return
	}

	for i, task := range tasks {
		mark := "[ ]"
		if task.Done {
			mark = "[x]"
		}

		fmt.Printf("%d. %s %s\n", i+1, mark, task.Title)
	}
}

func addTask(tasks []Task) []Task {
	fmt.Println("название задачи:")
	title := readLine()

	if title == "" {
		fmt.Println("название задачи не может быть пустым")
		return tasks
	}

	tasks = append(tasks, Task{Title: title})
	return tasks
}

func doneTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("список пуст")
		return tasks
	}

	showTasks(tasks)

	fmt.Println("введи номер задачи:")

	n, err := strconv.Atoi(readLine())
	if err != nil {
		fmt.Println("введи номер пункта цифрой")
		return tasks
	}

	if n < 1 || n > len(tasks) {
		fmt.Println("нет такого пункта")
		return tasks
	}

	i := n - 1
	tasks[i].Done = true
	fmt.Printf("задача [%d] выполнена\n", n)
	return tasks
}

func deleteTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("список пуст")
		return tasks
	}

	showTasks(tasks)

	fmt.Println("введи номер задачи:")

	n, err := strconv.Atoi(readLine())
	if err != nil {
		fmt.Println("введи номер пункта цифрой")
		return tasks
	}

	if n < 1 || n > len(tasks) {
		fmt.Println("нет такого пункта")
		return tasks
	}

	i := n - 1
	tasks = append(tasks[:i], tasks[i+1:]...)
	fmt.Printf("задача [%d] удалена\n", n)
	return tasks
}

func main() {
	fmt.Println("todo-лист")

	tasks := []Task{}

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. показать задачи")
		fmt.Println("2. добавить задачу")
		fmt.Println("3. отметить выполненной задачу")
		fmt.Println("4. удалить задачу")
		fmt.Println("0. выход")

		choice, err := strconv.Atoi(readLine())
		if err != nil {
			fmt.Println("введи номер пункта цифрой")
			continue
		}

		switch choice {
		case 1:
			showTasks(tasks)
		case 2:
			tasks = addTask(tasks)
		case 3:
			tasks = doneTask(tasks)
		case 4:
			tasks = deleteTask(tasks)
		case 0:
			return
		default:
			fmt.Println("нет такого пункта")
		}
	}
}
