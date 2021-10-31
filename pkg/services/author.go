package services

import (
	"basic-golang-api/pkg/db"
	m "basic-golang-api/pkg/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var input m.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addModel := m.Author{
		Name:        input.Name,
		Surname:     input.Surname,
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

func DeleteAuthor(c *gin.Context) {

	var input m.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteModel := m.Author{
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

func UpdateAuthor(c *gin.Context) {

	var input m.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateModel := m.Author{
		Id:          input.Id,
		Name:        input.Name,
		Surname:     input.Surname,
		UpdatedDate: time.Now(),
		UpdatedBy:   "App",
	}

	_, err := db.NewDBConn().Model(&updateModel).Where("id = ?", updateModel.Id).Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Update successful"})

}
