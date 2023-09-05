package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID         string `json:"id"`
	Task       string `json:"task"`
	IsComplete bool   `json:"isComplete"`
}

var tasks = []todo{
	{
		ID:         "1",
		Task:       "Learn the basics about golang",
		IsComplete: false,
	},
}

func main() {
	r := gin.Default()

	r.GET("/tasks", getTasks)
	r.POST("/tasks", addTask)
	r.GET("/task/:id", getOneTask)
	// r.PATCH("/task:id", updateOneTask)

	r.Run("localhost:4000")

}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var newTask todo

	if err := c.BindJSON(&newTask); err != nil {
		panic(err)
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getOneById(id string) (*todo, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("todo not found")
}

func getOneTask(c *gin.Context) {
	id := c.Param("id")

	t, err := getOneById(id)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, t)
}

// func updateOneTask(c *gin.Context) *todo {
// 	fmt.Println(c, "<<<<")

// 	id := c.Param("id")
// 	t, err := getOneById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
// 	}

// 	// var newTask todo
// 	// for i, v := range tasks {
// 	// 	if v.ID == id {
// 	// 		if err := c.BindJSON(&newTask); err != nil {
// 	// 			panic(err)
// 	// 		}

// 	// 	}
// 	// }
// }
