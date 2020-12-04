package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	CreatedAt       time.Time
	CreatedAtString string
	CreatedBy       string
	Content         string
	Status          int
}

func main() {
	//fmt.Println("hello wolrd!")
	todo := make([]Todo, 0, 5)

	t1 := Todo{
		CreatedBy:       "John",
		Content:         "go to the Gym",
		Status:          0,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t2 := Todo{
		CreatedBy:       "Mike",
		Content:         "watch movies",
		Status:          0,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t3 := Todo{
		CreatedBy:       "Harry",
		Content:         "go shopping",
		Status:          1,
		CreatedAt:       time.Now(),
		CreatedAtString: time.Now().Format("2006-01-02 15:04:05"),
	}
	t4 := Todo{
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
