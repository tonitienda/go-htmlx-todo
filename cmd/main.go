package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/tonitienda/go-htmlx-todo/pkg/tasks"
)

func index(w http.ResponseWriter, req *http.Request) {
	index := `<!doctype html>
	<html lang="en">
	<head>
		<script src="https://unpkg.com/htmx.org@1.9.8"></script>

		<!-- Latest compiled and minified CSS -->
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.1.3/dist/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous" />
		<!-- Optional theme -->
		<link rel="stylesheet" href="https://bootswatch.com/5/quartz/bootstrap.min.css" />
	</head>
	<body>
		<div class="main">
    		<h1>TODO</h1>
			<div class="col-sm-6" hx-get="/tasks" hx-trigger="load" />
		</div>
	</body>
	</html>`

	fmt.Fprint(w, index)
}

func getTasks(w http.ResponseWriter, req *http.Request) {
	tasksTemplate := `
	<form>
  <fieldset>
 	{{range $task := .}}

  		<div class="form-group row">
			<div class="col-sm-3">
				<input type="checkbox" hx-post="/tasks/{{ $task.ID }}/done" hx-trigger="changed"/>
			</div>
			<div class="col-sm-9">
				{{ $task.Title }}
			</div>
		</div>
	{{ end }}
	</fieldset>
	</form>`

	tasks := tasks.GetTasks()

	t := template.Must(template.New("Tasks").Parse(tasksTemplate))
	t.Execute(w, tasks)
}

func markTaskAsDone(w http.ResponseWriter, req *http.Request) {
	// Task ID  is in the URL
	//taskId := strings.TrimPrefix(req.URL.Path, "/tasks/")
	fmt.Fprint(w, "Task marked as done")

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/tasks/:id/done", markTaskAsDone)

	http.ListenAndServe(":8090", nil)
}
