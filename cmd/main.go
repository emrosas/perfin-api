package main

import (
	"log"

	"github.com/emrosas/perfin-api/pkg/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/create_table", handlers.CreateTableHandler)
	r.POST("/insert_data", handlers.InsertDataHandler)
	r.GET("/fetch_data", handlers.FetchDataHandler)
	r.Run() // listen and serve on 0.0.0.0:8080r := gin.Default()
}
