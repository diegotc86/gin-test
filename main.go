package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	Id           string `json:"id"`
	Body         string `json:"body"`
	Completed    bool   `json:"completed"`
	Created_date string `json:"created_date"`
}

var tasks = []task{
	{Id: "10", Body: "Task 1", Completed: false, Created_date: "2022-09-09T15:00:00+07:00"},
	{Id: "11", Body: "Task 2", Completed: true, Created_date: "2022-09-10T15:00:00+07:00"},
	{Id: "12", Body: "Task 3", Completed: true, Created_date: "2022-09-11T15:00:00+07:00"},
}

func getPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func indexTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(id string) (*task, error) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("Task not found")
}

func showTask(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func main() {
	router := gin.Default()

	router.GET("/ping", getPing)
	// get /tasks -> indexTasks
	router.GET("/tasks", indexTasks)
	// get /tasks/:id -> showTask
	router.GET("/tasks/:id", showTask)
	// post /tasks -> createTask body -> {"id": "13", "body": "New task", "completed": false, "created_date": "2022-09-12T00:00:00-05:00"} retorna Created(201) la tarea creada
	// patch /tasks/:id -> updateTask body -> {"completed": true} retorna OK(200) la tarea actualizada
	// delete /tasks/:id -> destroyTask retorna OK(200) NoContent

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
