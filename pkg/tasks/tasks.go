package tasks

type Task struct {
	ID    int
	Title string
	Done  bool
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
