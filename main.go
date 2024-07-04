package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

//`id` int NOT NULL AUTO_INCREMENT,
//`title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//`image` json DEFAULT NULL,
//`description` text,
//`status` enum('Doing','Done','Deleted') DEFAULT 'Doing',
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println("Connected to database", db)

	fmt.Println("Hello, World!")

	//now := time.Now()
	//
	//item := TodoItem{
	//	Id:          1,
	//	Title:       "Learn Golang",
	//	Description: "Learn Golang for beginners",
	//	Status:      "Doing",
	//	CreatedAt:   &now,
	//	UpdatedAt:   &now,
	//}
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var data TodoItemCreation
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var data TodoItem

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var data TodoItemUpdate

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	}
}
