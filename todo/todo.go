package todo

import (
	"errors"
	"fmt"
	"time"
)

type Storage interface {
	AddTask(title string, desc string) error
	ListTasks()
	DeleteTask(title string) error
	DoneTask(title string) error
	EditTask(title string, newDesc string) error
	AddEvent(input string, errText string)
	ListEvents()
}

type MemoryStorage struct {
	cases  map[string]Task
	events []Event
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		cases: make(map[string]Task),
	}
}

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

func (s *MemoryStorage) AddTask(title string, desc string) error {
	_, ok := s.cases[title]
	if ok {
		return errors.New("задача c таким заголовком уже есть")
	}

	s.cases[title] = Task{
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
		Done:        false,
	}
	return nil
}

func (s *MemoryStorage) ListTasks() {
	if len(s.cases) == 0 {
		fmt.Println("список задач пока пуст")
	} else {
		for k, task := range s.cases {
			fmt.Println(k, task)
		}
	}
}

func (s *MemoryStorage) DeleteTask(title string) error {
	_, ok := s.cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	} else {
		delete(s.cases, title)
	}
	return nil
}

func (s *MemoryStorage) DoneTask(title string) error {
	task, ok := s.cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	} else {
		now := time.Now()
		task.Done = true
		task.CompletedAt = &now
		s.cases[title] = task
	}
	return nil
}

func (s *MemoryStorage) EditTask(title string, newDesc string) error {
	task, ok := s.cases[title]
	if !ok {
		return errors.New("такой задачи нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	} else {
		task.Description = newDesc
		s.cases[title] = task
	}
	return nil
}

func (s *MemoryStorage) AddEvent(input string, errText string) {
	s.events = append(s.events, Event{input, errText, time.Now()})
}

func (s *MemoryStorage) ListEvents() {
	if len(s.events) == 0 {
		fmt.Println("список событий пуст!")
	} else {
		for _, event := range s.events {
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
