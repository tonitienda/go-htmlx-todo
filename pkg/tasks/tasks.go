package tasks

import (
	"time"

	"github.com/xeonx/timeago"
)

type Task struct {
	ID          int
	Title       string
	Description string
	AddedAt     time.Time
	DoneAt      time.Time
}

func (t Task) IsDone() bool {
	return !t.DoneAt.IsZero()
}

func (t Task) DoneTimeAgo() string {
	return timeago.English.FormatRelativeDuration(time.Since(t.DoneAt))
}

func (t Task) AddedTimeAgo() string {
	return timeago.English.FormatRelativeDuration(time.Since(t.AddedAt))
}

func (t Task) Took() string {
	return t.DoneAt.Sub(t.AddedAt).String()
}

var inMemoryTasks map[int]Task

func init() {
	inMemoryTasks = map[int]Task{
		1: {
			ID:          1,
			Title:       "First task",
			Description: "This is the description for the first task",
			AddedAt:     time.Date(2023, 11, 05, 8, 30, 0, 0, time.UTC),
			DoneAt:      time.Date(2023, 11, 06, 10, 20, 0, 0, time.UTC),
		},
		2: {
			ID:          2,
			Title:       "Second task",
			Description: "This is the description for the second task",
			AddedAt:     time.Date(2023, 11, 04, 8, 30, 0, 0, time.UTC),
		},
	}
}

func GetTasks() []Task {
	// Convert map to slice of values
	tasks := make([]Task, 0, len(inMemoryTasks))
	for _, task := range inMemoryTasks {
		tasks = append(tasks, task)
	}

	return tasks

}

func AddTask(title string, description string) (int, error) {
	// TODO - Add validation

	id := len(inMemoryTasks) + 1
	inMemoryTasks[id] = Task{
		ID:          id,
		Title:       title,
		Description: description,
		AddedAt:     time.Now(),
	}
	return id, nil
}

func MarkAsDone(id int) {
	// Convert map to slice of values
	task, ok := inMemoryTasks[id]
	if !ok {
		return
	}

	task.DoneAt = time.Now()
	inMemoryTasks[id] = task
}

func MarkAsTodo(id int) {
	// Convert map to slice of values
	task, ok := inMemoryTasks[id]
	if !ok {
		return
	}

	// TODO - See how to reset DoneAt
	// 	task.DoneAt = nil

	newTask := Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		AddedAt:     task.AddedAt,
	}
	inMemoryTasks[id] = newTask
}
