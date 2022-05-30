package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/models"
	"github.com/fredmessias43/rest-api/utils"
)

type MailController struct {
	DB *gorm.DB
}

func (p *MailController) Index(c *gin.Context) {
	mails := []models.Mail{}
	p.DB.Find(&mails)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   mails,
	})
}

func (p *MailController) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	mail := models.Mail{}
	_ = p.DB.Find(&mail, ID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   mail,
	})
}

func (p *MailController) Store(c *gin.Context) {
	var mail models.Mail

	if err := c.ShouldBindJSON(&mail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"messages": utils.Format422Error(err.Error()),
		})
		return
	}

	_ = p.DB.Create(&mail)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"data":   mail,
	})
}

func (p *MailController) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	mail := models.Mail{}
	_ = p.DB.Find(&mail, ID)

	if err := c.ShouldBindJSON(&mail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"messages": utils.Format422Error(err.Error()),
		})
		return
	}

	dbResult := p.DB.Model(&mail).Where("ID = ?", ID).Updates(&mail)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"messages": dbResult.Error,
		})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": http.StatusAccepted,
		"data":   mail,
	})
}

func (p *MailController) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	mail := models.Mail{}
	_ = p.DB.Delete(&mail, int(ID))

	c.JSON(http.StatusNoContent, "")
}
