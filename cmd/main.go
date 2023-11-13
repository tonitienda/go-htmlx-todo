package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"github.com/gin-gonic/gin"
	"github.com/tonitienda/go-htmlx-todo/pkg/tasks"
)

func index(c *gin.Context) {
	fmt.Println("Index")

	tasks := tasks.GetTasks()

	templData := map[string]interface{}{
		"component": "task-mgmt.tmpl",
		"data":      tasks,
	}

	c.HTML(http.StatusOK, "index.tmpl", templData)
}

func getTasks(c *gin.Context) {
	fmt.Println("getTasks")

	tasks := tasks.GetTasks()

	c.HTML(http.StatusOK, "tasks.tmpl", tasks)
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
	//router := gin.Default()

	// Pre compile templates
	dat, err := os.ReadFile("templates/index.tmpl")

	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
	contents := strings.ReplaceAll(string(dat), "### COMPONENT ###", "task-mgmt.tmpl")

	fmt.Print(contents)

	t1 := template.New("test")

	fmt.Print(t1)
	t1 = template.Must(t1.Parse(contents))

	fmt.Print(t1)

	// router.LoadHTMLGlob("templates/*")

	// router.GET("/", index)
	// router.GET("/tasks", getTasks)
	// router.POST("/tasks", addTask)
	// router.POST("/tasks/:id/done", markTaskAsDone)
	// router.POST("/tasks/:id/todo", markAsTodo)

	// router.Run(":8080")
}
