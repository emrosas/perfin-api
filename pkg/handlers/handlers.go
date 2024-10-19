package handlers

import (
	"net/http"

	"github.com/emrosas/perfin-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func CreateTableHandler(c *gin.Context) {
	conn, err := database.InitDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer conn.Close(c.Request.Context())

	err = database.CreateTable(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Table created successfully",
	})
}

func InsertDataHandler(c *gin.Context) {
	conn, err := database.InitDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer conn.Close(c.Request.Context())

	err = database.InsertData(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data inserted successfully",
	})
}

func FetchDataHandler(c *gin.Context) {
	conn, err := database.InitDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer conn.Close(c.Request.Context())

	data, err := database.QueryData(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
