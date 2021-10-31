package main

import (
	m "basic-golang-api/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/createbook", m.CreateBook)
	router.POST("/bookbyid", m.GetBookById)
	router.GET("/allbooks", m.AllBooks)
	router.PUT("/updatebook", m.UpdateBook)
	router.DELETE("/deletebook", m.DeleteBook)

	router.Run("localhost:9093")
}
