package todo

import (
	"errors"
	"fmt"
	"time"
)

var events []Event
var cases = make(map[string]Task)

type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	CompletedAt *time.Time
	Done        bool
}

type Event struct {
	Input     string
	ErrorText string
	CreatedAt time.Time
}

func AddTask(title string, desc string) error {
	_, ok := cases[title]
	if ok {
		return errors.New("задача c таким заголовком уже есть")
	}

	cases[title] = Task{
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
		Done:        false,
	}
	return nil
}

func ListTasks() {
	if len(cases) == 0 {
		fmt.Println("список задач пока пуст")
	} else {
		for k, task := range cases {
			fmt.Println(k, task)
		}
	}
}

func DeleteTask(title string) error {
	_, ok := cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	} else {
		delete(cases, title)
	}
	return nil
}

func DoneTask(title string) error {
	task, ok := cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	} else {
		now := time.Now()
		task.Done = true
		task.CompletedAt = &now
		cases[title] = task
	}
	return nil
}

func EditTask(title string, newDesc string) error {
	task, ok := cases[title]
	if !ok {
		return errors.New("такой задачи нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	} else {
		task.Description = newDesc
		cases[title] = task
	}
	return nil
}

func AddEvent(input string, errText string) {
	events = append(events, Event{input, errText, time.Now()})
}

func ListEvents() {
	if len(events) == 0 {
		fmt.Println("список событий пуст!")
	} else {
		for _, event := range events {
			fmt.Println(event)
		}
	}
}

func PrintHelp() {
	fmt.Println(`
Доступные команды:
- Помощь — список всех команд и их формат
- Добавить {заголовок} {текст задачи} — добавить новую задачу
- Список — показать все задачи
- Изменить - {заголовок} {новый текст} изменяет существующую задачу
- Удалить {заголовок} — удалить задачу по заголовку
- Выполнить {заголовок} — отметить задачу выполненной
- События — показать историю всех введённых команд
- Выход — завершить работу программы`)
}
