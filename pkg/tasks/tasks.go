package tasks

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

var inMemoryTasks map[int]Task

func init() {
	inMemoryTasks = map[int]Task{
		1: {
			ID:    1,
			Title: "First task",
			Done:  true,
		},
		2: {
			ID:    2,
			Title: "Second task",
			Done:  false,
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
		Done:        false,
	}
	return id, nil
}

func MarkAsDone(id int) {
	// Convert map to slice of values
	task, ok := inMemoryTasks[id]
	if !ok {
		return
	}

	task.Done = true
	inMemoryTasks[id] = task
}

func MarkAsTodo(id int) {
	// Convert map to slice of values
	task, ok := inMemoryTasks[id]
	if !ok {
		return
	}

	task.Done = false
	inMemoryTasks[id] = task
}
