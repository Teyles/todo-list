package todo

import (
	"fmt"
	"strings"
)

func HandlerCommand(text string) bool {

	fields := strings.Fields(text)

	if len(fields) == 0 {
		AddEvent(text, "Пустой ввод")
		return false
	}

	errText := ""

	switch strings.ToLower(fields[0]) {
	case "добавить":
		if len(fields) < 3 {
			errText = "He хватает заголовка или текста"
		} else {
			desc := strings.Join(fields[2:], " ")
			errText = AddTask(fields[1], desc)
			if errText == "" {
				fmt.Println("Задача добавлена!")
			}
		}
	case "список":
		if len(fields) == 1 {
			ListTasks()
		} else {
			errText = "Лишние аргументы y команды"
		}
	case "удалить":
		if len(fields) != 2 {
			errText = "Неверное количество слов"
		} else {
			errText = DeleteTask(fields[1])
			if errText == "" {
				fmt.Println("Задача удалена!")
			}
		}
	case "выполнить":
		if len(fields) != 2 {
			errText = "Неправильно введена команда!"
		} else {
			errText = DoneTask(fields[1])
			if errText == "" {
				fmt.Println("Задача выполнена!")
			}
		}
	case "изменить":
		if len(fields) < 3 {
			errText = "Неверно введена команда!"
		} else {
			newDesc := strings.Join(fields[2:], " ")
			errText = EditTask(fields[1], newDesc)
			if errText == "" {
				fmt.Println("Задача изменена!")
			}
		}
	case "события":
		if len(fields) == 1 {
			ListEvents()
		} else {
			errText = "Лишние аргументы у команды"
		}
	case "помощь":
		if len(fields) == 1 {
			PrintHelp()
		} else {
			errText = "Лишние аргументы у команды"
		}
	case "выход":
		if len(fields) == 1 {
			AddEvent(text, "")
			return true
		} else {
			errText = "Лишние аргументы у команды"
		}
	default:
		errText = "Неизвестная команда"
	}

	if errText != "" {
		fmt.Println("Ошибка: ", errText)
	}
	AddEvent(text, errText)
	return false
}
