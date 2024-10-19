package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/your-username/financial-api-app/pkg/database"
	"github.com/your-username/financial-api-app/pkg/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/create_table", handlers.CreateTableHandler)
	r.POST("/insert_data", handlers.InsertDataHandler)
	r.GET("/query_data", handlers.QueryDataHandler)
	r.Run() // listen and serve on 0.0.0.0:8080r := gin.Default()
}
