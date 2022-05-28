package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/models"
)

type PersonController struct {
	DB *gorm.DB
}

func (p *PersonController) Index(c *gin.Context) {
	people := []models.Person{}
	p.DB.Find(&people)
	c.JSON(200, people)
}

func (p *PersonController) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	person := models.Person{}
	_ = p.DB.Find(&person, int(ID))

	c.JSON(200, person)
}

func (p *PersonController) Store(c *gin.Context) {
	var person models.Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_ = p.DB.Create(&person)

	c.JSON(201, person)
}

func (p *PersonController) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	person := models.Person{}
	_ = p.DB.Delete(&person, int(ID))

	c.JSON(200, person)
}
