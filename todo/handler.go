package todo

import (
	"errors"
	"fmt"
	"strings"
)

func HandlerCommand(text string) bool {

	fields := strings.Fields(text)

	if len(fields) == 0 {
		AddEvent(text, "Пустой ввод")
		return false
	}

	var err error

	switch strings.ToLower(fields[0]) {
	case "добавить":
		if len(fields) < 3 {
			err = errors.New("нe хватает заголовка или текста")
		} else {
			desc := strings.Join(fields[2:], " ")
			err = AddTask(fields[1], desc)
			if err == nil {
				fmt.Println("Задача добавлена!")
			}
		}
	case "список":
		if len(fields) == 1 {
			ListTasks()
		} else {
			err = errors.New("лишние аргументы y команды")
		}
	case "удалить":
		if len(fields) != 2 {
			err = errors.New("неверное количество слов")
		} else {
			err = DeleteTask(fields[1])
			if err == nil {
				fmt.Println("Задача удалена!")
			}
		}
	case "выполнить":
		if len(fields) != 2 {
			err = errors.New("неправильно введена команда")
		} else {
			err = DoneTask(fields[1])
			if err == nil {
				fmt.Println("Задача выполнена!")
			}
		}
	case "изменить":
		if len(fields) < 3 {
			err = errors.New("неверно введена команда")
		} else {
			newDesc := strings.Join(fields[2:], " ")
			err = EditTask(fields[1], newDesc)
			if err == nil {
				fmt.Println("Задача изменена!")
			}
		}
	case "события":
		if len(fields) == 1 {
			ListEvents()
		} else {
			err = errors.New("лишние аргументы у команды")
		}
	case "помощь":
		if len(fields) == 1 {
			PrintHelp()
		} else {
			err = errors.New("лишние аргументы у команды")
		}
	case "выход":
		if len(fields) == 1 {
			AddEvent(text, "")
			return true
		} else {
			err = errors.New("Лишние аргументы у команды")
		}
	default:
		err = errors.New("Неизвестная команда")
	}

	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		AddEvent(text, err.Error())
	} else {
		AddEvent(text, "")
	}
	return false
}
