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

func getTaskById(id string) (*task, error, int) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			return &tasks[i], nil, i
		}
	}
	return nil, errors.New("Task not found"), 0
}

func showTask(c *gin.Context) {
	id := c.Param("id")
	task, err, _ := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}
func createTask(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
func updateTask(c *gin.Context) {
	// id := c.Param("id")
	var editTask task
	if err := c.BindJSON(&editTask); err != nil {
		return
	}
	//update
	c.JSON(200, editTask)
}
func destroyTask(c *gin.Context) {
	id := c.Param("id")
	task, err, indx := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks[:indx], tasks[indx+1:]...)
	c.IndentedJSON(http.StatusOK, task)
}

func main() {
	router := gin.Default()

	router.GET("/ping", getPing)
	// get /tasks -> indexTasks
	router.GET("/tasks", indexTasks)
	// get /tasks/:id -> showTask
	router.GET("/tasks/:id", showTask)
	// post /tasks -> createTask bodty -> {"id": "13", "body": "New task", "completed": false, "created_date": "2022-09-12T00:00:00-05:00"} retorna Created(201) la tarea creada
	router.POST("/tasks", createTask)
	// patch /tasks/:id -> updateTask body -> {"completed": true} retorna OK(200) la tarea actualizada
	router.PATCH("/tasks/:id", updateTask)
	// delete /tasks/:id -> destroyTask retorna OK(200) NoContent
	router.DELETE("/tasks/:id", destroyTask)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
