package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID              int `form:"id"`
	CreatedAt       time.Time
	CreatedAtString string
	CreatedBy       string `form:"createdBy"`
	Content         string `form:"content"`
	Status          int    `form:"status"`
}

var todo []Todo
var idMax = 0

func incremetId() int {
	idMax = idMax + 1
	return idMax
}

func GetDataTodo(c *gin.Context) {
	var t Todo
	c.Bind(&t)
	t.ID = incremetId()
	t.Status = 0
	t.CreatedAtString = time.Now().Format("2006-01-02 15:04:05")
	todo = append(todo, t)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func GetDoneTodo(c *gin.Context) {
	var t Todo
	c.Bind(&t)

	if t.Status == 0 {
		t.Status = 1
	} else {
		t.Status = 0
	}

	for idx, tt := range todo {
		if tt.ID == t.ID {
			todo[idx].Status = t.Status
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func main() {
	//fmt.Println("hello wolrd!")
	todo = make([]Todo, 0, 5)

	t1 := Todo{
		ID:              incremetId(),
		CreatedBy:       "John",
		Content:         "go to the Gym",
		Status:          0,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t2 := Todo{
		ID:              incremetId(),
		CreatedBy:       "Mike",
		Content:         "watch movies",
		Status:          0,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t3 := Todo{
		ID:              incremetId(),
		CreatedBy:       "Harry",
		Content:         "go shopping",
		Status:          1,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t4 := Todo{
		ID:              incremetId(),
		CreatedBy:       "Steve",
		Content:         "wash a car",
		Status:          0,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	todo = append(todo, t1)
	todo = append(todo, t2)
	todo = append(todo, t3)
	todo = append(todo, t4)

	r := gin.Default()
	r.Static("/template", "./template")
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wolrd",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todo,
		})
	})
	r.GET("/create", GetDataTodo)
	r.GET("/done", GetDoneTodo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
