package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/go-htmlx-todo/pkg/tasks"
)

func index(c *gin.Context) {
	fmt.Println("Index")

	tasks := tasks.GetTasks()
	fmt.Println(tasks)
	c.HTML(http.StatusOK, "task-mgmt.tmpl", tasks)
}

func getTasks(c *gin.Context) {
	fmt.Println("getTasks")

	tasks := tasks.GetTasks()

	c.HTML(http.StatusOK, "components/tasks.tmpl", tasks)
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
	router.LoadHTMLGlob("templates/**/*.tmpl")

	router.GET("/", index)
	router.GET("/tasks", getTasks)
	router.POST("/tasks", addTask)
	router.POST("/tasks/:id/done", markTaskAsDone)
	router.POST("/tasks/:id/todo", markAsTodo)

	router.Run(":8080")
}
