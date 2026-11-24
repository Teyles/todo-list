// Package todo содержит логику хранения задач и обработки пользовательских команд.
package todo

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Storage определяет базовые методы для работы с задачами и историей.
type Storage interface {
	AddTask(title string, desc string) error
	ListTasks()
	DeleteTask(title string) error
	DoneTask(title string) error
	EditTask(title string, newDesc string) error
	AddEvent(input string, errText string)
	ListEvents()
}

// MemoryStorage хранит все задачи и события в оперативной памяти.
type MemoryStorage struct {
	cases  map[string]Task
	events []Event
	mu     sync.Mutex
}

// NewMemoryStorage инициализирует и возвращает новое хранилище.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		cases:  make(map[string]Task),
		events: make([]Event, 0),
	}
}

// Task описывает структуру одной задачи.
type Task struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Title       string
	Description string
	Done        bool
}

// Event описывает структуру одного события в истории команд.
type Event struct {
	CreatedAt time.Time
	Input     string
	ErrorText string
}

// AddTask добавляет новую задачу с указанным заголовком и описанием.
func (s *MemoryStorage) AddTask(title, desc string) error {
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

// ListTasks выводит все текущие задачи в терминал.
func (s *MemoryStorage) ListTasks() {
	if len(s.cases) == 0 {
		fmt.Println("список задач пока пуст")
	} else {
		for title, task := range s.cases {
			status := "[ ]"
			if task.Done {
				status = "[X]"
			}
			fmt.Printf("%s %s — %s\n", status, title, task.Description)
		}
	}
}

// EditTask обновляет текстовое описание для существующей задачи.
func (s *MemoryStorage) EditTask(title, newDesc string) error {
	task, ok := s.cases[title]
	if !ok {
		return errors.New("такой задачи нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	}
	task.Description = newDesc
	s.cases[title] = task
	return nil
}

// DeleteTask удаляет задачу из хранилища по её заголовку.
func (s *MemoryStorage) DeleteTask(title string) error {
	_, ok := s.cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	}
	delete(s.cases, title)
	return nil
}

// DoneTask отмечает указанную задачу как выполненную и фиксирует время.
func (s *MemoryStorage) DoneTask(title string) error {
	task, ok := s.cases[title]
	if !ok {
		return errors.New("задачи c таким заголовком нет")
	} else if task.Done {
		return errors.New("задача уже выполнена")
	}
	task.Done = true
	task.CompletedAt = time.Now()
	s.cases[title] = task
	return nil
}

// AddEvent сохраняет введенную команду и текст ошибки в историю.
func (s *MemoryStorage) AddEvent(input, errText string) {
	s.events = append(s.events, Event{
		CreatedAt: time.Now(),
		Input:     input,
		ErrorText: errText,
	})
}

// ListEvents выводит историю всех обработанных команд.
func (s *MemoryStorage) ListEvents() {
	if len(s.events) == 0 {
		fmt.Println("список событий пуст!")
	} else {
		for _, event := range s.events {
			fmt.Printf("[%s] Команда: %s, Ошибка: %s\n", event.CreatedAt.Format("15:04:05"), event.Input, event.ErrorText)
		}
	}
}

// PrintHelp показывает пользователю список доступных команд.
func PrintHelp() {
	fmt.Println(`
Доступные команды:
- Добавить {заголовок} {текст задачи} — добавить новую задачу
- Список — показать все задачи
- Изменить - {заголовок} {новый текст} изменяет существующую задачу
- Удалить {заголовок} — удалить задачу по заголовку
- Выполнить {заголовок} — отметить задачу выполненной
- События — показать историю всех введённых команд
- Помощь — список всех команд и их формат
- Выход — завершить работу программы`)
}
