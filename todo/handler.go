package todo

import (
	"errors"
	"fmt"
	"strings"
)

func HandlerCommand(text string, store Storage) bool {

	fields := strings.Fields(text)

	if len(fields) == 0 {
		store.AddEvent(text, "Пустой ввод")
		return false
	}

	var err error

	switch strings.ToLower(fields[0]) {
	case "добавить":
		if len(fields) < 3 {
			err = errors.New("нe хватает заголовка или текста")
		} else {
			desc := strings.Join(fields[2:], " ")
			err = store.AddTask(fields[1], desc)
			if err == nil {
				fmt.Println("Задача добавлена!")
			}
		}
	case "список":
		if len(fields) == 1 {
			store.ListTasks()
		} else {
			err = errors.New("лишние аргументы y команды")
		}
	case "удалить":
		if len(fields) != 2 {
			err = errors.New("неверное количество слов")
		} else {
			err = store.DeleteTask(fields[1])
			if err == nil {
				fmt.Println("Задача удалена!")
			}
		}
	case "выполнить":
		if len(fields) != 2 {
			err = errors.New("неправильно введена команда")
		} else {
			err = store.DoneTask(fields[1])
			if err == nil {
				fmt.Println("Задача выполнена!")
			}
		}
	case "изменить":
		if len(fields) < 3 {
			err = errors.New("неверно введена команда")
		} else {
			newDesc := strings.Join(fields[2:], " ")
			err = store.EditTask(fields[1], newDesc)
			if err == nil {
				fmt.Println("Задача изменена!")
			}
		}
	case "события":
		if len(fields) == 1 {
			store.ListEvents()
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
			store.AddEvent(text, "")
			return true
		} else {
			err = errors.New("Лишние аргументы у команды")
		}
	default:
		err = errors.New("Неизвестная команда")
	}

	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		store.AddEvent(text, err.Error())
	} else {
		store.AddEvent(text, "")
	}
	return false
}
