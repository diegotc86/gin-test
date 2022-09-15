package tasks

import (
	"errors"
	"net/http"

	"github.com/diegotc86/gin-test/bd"
	"github.com/gin-gonic/gin"
)

func getTaskById(id string) (*bd.Task, error, int) {
	for i := 0; i < len(bd.Tasks); i++ {
		if bd.Tasks[i].Id == id {
			return &bd.Tasks[i], nil, i
		}
	}
	return nil, errors.New("Task not found"), 0
}

func IndexTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bd.Tasks)
}

func ShowTask(c *gin.Context) {
	id := c.Param("id")
	task, err, _ := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask bd.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	bd.Tasks = append(bd.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	task, err, _ := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		return
	}

	//update
	c.JSON(200, task)
}

func DestroyTask(c *gin.Context) {
	id := c.Param("id")
	task, err, indx := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	bd.Tasks = append(bd.Tasks[:indx], bd.Tasks[indx+1:]...)

	c.IndentedJSON(http.StatusOK, task)
}
