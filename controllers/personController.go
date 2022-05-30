package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/models"
	"github.com/fredmessias43/rest-api/utils"
)

type PersonController struct {
	DB *gorm.DB
}

func (p *PersonController) Index(c *gin.Context) {
	people := []models.Person{}
	p.DB.Find(&people)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   people,
	})
}

func (p *PersonController) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	person := models.Person{}
	_ = p.DB.Find(&person, ID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   person,
	})
}

func (p *PersonController) Store(c *gin.Context) {
	var person models.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"messages": utils.Format422Error(err.Error()),
		})
		return
	}

	_ = p.DB.Create(&person)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"data":   person,
	})
}

func (p *PersonController) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	person := models.Person{}
	_ = p.DB.Find(&person, ID)

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"messages": utils.Format422Error(err.Error()),
		})
		return
	}

	dbResult := p.DB.Model(&person).Where("ID = ?", ID).Updates(&person)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"messages": dbResult.Error,
		})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": http.StatusAccepted,
		"data":   person,
	})
}

func (p *PersonController) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	person := models.Person{}
	_ = p.DB.Delete(&person, int(ID))

	c.JSON(http.StatusNoContent, "")
}
