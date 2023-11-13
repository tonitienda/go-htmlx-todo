package tasks

import (
	"time"

	"github.com/google/uuid"
	"github.com/xeonx/timeago"
)

type Task struct {
	ID          string
	Title       string
	Description string
	AddedAt     time.Time
	DoneAt      time.Time
	DependsOn   string
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

var inMemoryTasks map[string]Task

func init() {
	inMemoryTasks = map[string]Task{
		"89c2553e-3f12-49a3-b429-3c9c15b76341": {
			ID:          "89c2553e-3f12-49a3-b429-3c9c15b76341",
			Title:       "First task",
			Description: "This is the description for the first task",
			AddedAt:     time.Date(2023, 11, 05, 8, 30, 0, 0, time.UTC),
			DoneAt:      time.Date(2023, 11, 06, 10, 20, 0, 0, time.UTC),
		},
		"aef9c208-0874-45af-94dc-87a48a2cd171": {
			ID:          "aef9c208-0874-45af-94dc-87a48a2cd171",
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

func AddTask(title string, description string, dependsOn string) (string, error) {
	// TODO - Add validation

	id := uuid.NewString()
	inMemoryTasks[id] = Task{
		ID:          id,
		Title:       title,
		Description: description,
		AddedAt:     time.Now(),
		DependsOn:   dependsOn,
	}
	return id, nil
}

func MarkAsDone(id string) {
	// Convert map to slice of values
	task, ok := inMemoryTasks[id]
	if !ok {
		return
	}

	task.DoneAt = time.Now()
	inMemoryTasks[id] = task
}

func MarkAsTodo(id string) {
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
