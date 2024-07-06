package main

import (
	"Learn_Golang/common"
	"Learn_Golang/modules/item/model"
	"Learn_Golang/modules/item/transport/ginitem"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

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
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
