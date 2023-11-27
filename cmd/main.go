package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/go-htmlx-todo/pkg/tasks"
)

type CounterCardData struct {
	Title string
	Count int
	Level string
}

func filter(allTasks []tasks.Task, predicate func(tasks.Task) bool) []tasks.Task {
	filtered := make([]tasks.Task, 0, len(allTasks))
	for _, task := range allTasks {
		if predicate(task) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func index(c *gin.Context) {
	fmt.Println("Index")

	allTasks := tasks.GetTasks()

	doneTasks := filter(allTasks, func(task tasks.Task) bool { return task.IsDone() })

	c.HTML(http.StatusOK, "index", gin.H{
		"Page":  "home",
		"Tasks": allTasks,
		"Cards": []CounterCardData{
			{
				Title: "All Tasks",
				Count: len(allTasks),
				Level: "info",
			},
			{
				Title: "Done",
				Count: len(doneTasks),
				Level: "success",
			},
		},
	})
}

func tasksTree(c *gin.Context) {
	fmt.Println("Tasks Tree")

	tasks := tasks.GetTasks()

	fmt.Println(tasks)
	c.HTML(http.StatusOK, "index", gin.H{
		"Page":  "tasks-tree",
		"Tasks": tasks,
	})
}

func tasksMgmt(c *gin.Context) {
	fmt.Println("Tasks Mgmt")

	tasks := tasks.GetTasks()

	fmt.Println(tasks)
	c.HTML(http.StatusOK, "index", gin.H{
		"Page":  "tasks-mgmt",
		"Tasks": tasks,
	})
}

func getTasks(c *gin.Context) {
	fmt.Println("getTasks")

	tasks := tasks.GetTasks()

	c.HTML(http.StatusOK, "components/tasks.html", tasks)
}

func markTaskAsDone(c *gin.Context) {
	// Task ID  is in the URL
	taskId := c.Param("id")

	tasks.MarkAsDone(taskId)
	getTasks(c)

}

func addTask(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")

	dependsOn := c.PostForm("dependsOn")

	tasks.AddTask(title, description, dependsOn)

	getTasks(c)
}

func markAsTodo(c *gin.Context) {
	// Task ID  is in the URL
	taskId := c.Param("id")

	tasks.MarkAsTodo(taskId)
	getTasks(c)

}

func main() {
	router := gin.Default()

	// Pre compile templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Full pages
	router.GET("/", index)
	router.GET("/tasks/tree", tasksTree)
	router.GET("/tasks/mgmt", tasksMgmt)

	// Fragments
	router.GET("/tasks", getTasks)

	// Commands
	router.POST("/tasks", addTask)
	router.POST("/tasks/:id/done", markTaskAsDone)
	router.POST("/tasks/:id/todo", markAsTodo)

	router.Run(":8080")
}
