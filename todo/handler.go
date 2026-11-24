package todo

import (
	"errors"
	"fmt"
	"strings"
)

const minFieldsForEdit = 3

var err error

func handleAdd(fields []string, store Storage) error {
	if len(fields) < minFieldsForEdit {
		return errors.New("недостаточно аргументов. Использование: Добавить {заголовок} {текст}")
	}
	title := fields[1]
	desc := strings.Join(fields[2:], " ")
	return store.AddTask(title, desc)
}

func handleDelete(fields []string, store Storage) error {
	if len(fields) < 2 {
		return errors.New("недостаточно аргументов. Использование: Удалить {заголовок}")
	}
	title := fields[1]
	return store.DeleteTask(title)
}

func handleDone(fields []string, store Storage) error {
	if len(fields) != 2 {
		return errors.New("неправильное количество аргументов. Использование: Выполнить {заголовок}")
	}
	err = store.DoneTask(fields[1])
	if err != nil {
		return err
	}
	fmt.Println("Задача выполнена!")
	return nil
}

func handleEdit(fields []string, store Storage) error {
	if len(fields) < minFieldsForEdit {
		return errors.New("недостаточно аргументов. Использование: Изменить {заголовок} {текст}")
	}
	newDesc := strings.Join(fields[2:], " ")
	return store.EditTask(fields[1], newDesc)
}

// HandlerCommand считывает введенную строку и вызывает нужный метод хранилища.
//
//nolint:cyclop // Это роутер команд, большой switch здесь необходим
func HandlerCommand(text string, store Storage) bool {
	fields := strings.Fields(text)

	if len(fields) == 0 {
		store.AddEvent(text, "Пустой ввод")
		return false
	}

	command := strings.ToLower(fields[0])

	switch command {
	case "добавить":
		err = handleAdd(fields, store)
	case "список":
		store.ListTasks()
	case "удалить":
		err = handleDelete(fields, store)
	case "выполнить":
		err = handleDone(fields, store)
	case "изменить":
		err = handleEdit(fields, store)
	case "события":
		if len(fields) == 1 {
			store.ListEvents()
		} else {
			err = errors.New("неправильное количество аргументов. Использование: События")
		}
	case "помощь":
		if len(fields) == 1 {
			PrintHelp()
		} else {
			err = errors.New("неправильное количество аргументов. Использование: Помощь")
		}
	case "выход":
		return false
	default:
		err = errors.New("неизвестная команда")
	}

	if command != "выход" && command != "список" {
		errText := ""
		if err != nil {
			errText = err.Error()
			fmt.Println("Ошибка:", errText)
		}
		store.AddEvent(text, errText)
	}
	return true
}
