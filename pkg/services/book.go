package services

import (
	"basic-golang-api/pkg/db"
	m "basic-golang-api/pkg/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {

	var input m.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addModel := m.Book{
		Isbn:        input.Isbn,
		Name:        input.Name,
		Author:      input.Author,
		Year:        input.Year,
		CreatedDate: time.Now(),
		CreatedBy:   "App",
		IsActive:    true,
	}

	_, err := db.NewDBConn().Model(&addModel).Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": addModel})
}

func DeleteBook(c *gin.Context) {

	var input m.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteModel := m.Book{
		Id:          input.Id,
		UpdatedDate: time.Now(),
		UpdatedBy:   "App",
		IsActive:    false,
	}

	_, err := db.NewDBConn().Model(&deleteModel).Where("id = ?", deleteModel.Id).Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"result": "Delete successful"})

}

func UpdateBook(c *gin.Context) {

	var input m.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateModel := &m.Book{
		Id:          input.Id,
		Isbn:        input.Isbn,
		Name:        input.Name,
		Author:      input.Author,
		Year:        input.Year,
		UpdatedDate: time.Now(),
		UpdatedBy:   "App",
	}

	_, err := db.NewDBConn().Model(&updateModel).WherePK().Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Update successful"})

}

func AllBooks(c *gin.Context) {

	var resultModel []m.Book
	err := db.NewDBConn().Model(&resultModel).Select()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": resultModel})
}

func GetBookById(c *gin.Context) {
	var input m.Book

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.NewDBConn().Model(&input).Where("id = ?", input.Id).Select()

	if err != nil {
		panic(err)
		//c.JSON(http.StatusBadRequest, gin.H{"result": err})
		//return
	}

	c.JSON(http.StatusOK, gin.H{"result": input})
}
