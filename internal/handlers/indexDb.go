package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/indexing/internal/storage"
)

type Request struct {
	IndexName  string `json:"index_name"`
	ColumnName string `json:"column_name"`
}

func CreateIndexHandler(c *gin.Context, db *sql.DB) {
	var request Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	executionTime, err := storage.IndexDb(db, request.IndexName, request.ColumnName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Index created successfully", "execution_time": executionTime})
}
