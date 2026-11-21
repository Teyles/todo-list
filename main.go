package main

import (
	"bufio"
	"fmt"
	"os"
	"todo/todo-list/todo"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	todo.PrintHelp()

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода: ", scanner.Err())
			return
		}

		text := scanner.Text()

		if todo.HandlerCommand(text) {
			return
		}
	}
}
