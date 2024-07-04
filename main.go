package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

//`id` int NOT NULL AUTO_INCREMENT,
//`title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//`image` json DEFAULT NULL,
//`description` text,
//`status` enum('Doing','Done','Deleted') DEFAULT 'Doing',
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

func main() {
	fmt.Println("Hello, World!")

	now := time.Now()

	item := TodoItem{
		Id:          1,
		Title:       "Learn Golang",
		Description: "Learn Golang for beginners",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsonDate, err := json.Marshal(item)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("JSON: ", string(jsonDate))

	jsonStr := `{"id":1,"title":"Learn Golang","description":"Learn Golang for beginners","status":"Doing","created_at":"2024-07-04T09:32:32.4265897+07:00","updated_at":null}`

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Item2: ", item2)
}
