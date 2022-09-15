package main

import (
	"github.com/diegotc86/gin-test/tasks"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// get /tasks -> indexTasks
	router.GET("/tasks", tasks.IndexTasks)
	// get /tasks/:id -> showTask
	router.GET("/tasks/:id", tasks.ShowTask)
	// post /tasks -> createTask bodty -> {"id": "13", "body": "New task", "completed": false, "created_date": "2022-09-12T00:00:00-05:00"} retorna Created(201) la tarea creada
	router.POST("/tasks", tasks.CreateTask)
	// patch /tasks/:id -> updateTask body -> {"completed": true} retorna OK(200) la tarea actualizada
	router.PATCH("/tasks/:id", tasks.UpdateTask)
	// delete /tasks/:id -> destroyTask retorna OK(200) NoContent
	router.DELETE("/tasks/:id", tasks.DestroyTask)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
