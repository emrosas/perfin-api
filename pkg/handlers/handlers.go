package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-username/financial-api-app/pkg/database"
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
