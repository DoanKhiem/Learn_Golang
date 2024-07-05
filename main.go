package main

import (
	"Learn_Golang/common"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return allItemStatuses[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid status")

}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan data from sql:", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprint("Failed to parse status from sql:", value))
	}
	*item = v

	return nil

}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(`"` + item.String() + `"`), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `"`, "")
	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemValue

	return nil
}

type TodoItem struct {
	common.SQLModel
	//Id          int         `json:"id" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
	//CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at"`
	//UpdatedAt   *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
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
			items.GET("", ListItem(db))
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Deleted",
		}); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		var result []TodoItem

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
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
