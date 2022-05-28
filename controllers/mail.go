package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/models"
)

type MailController struct {
	DB *gorm.DB
}

func (p *MailController) Index(c *gin.Context) {
	mails := []models.Mail{}
	p.DB.Find(&mails)
	c.JSON(200, mails)
}

func (p *MailController) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	mail := models.Mail{}
	_ = p.DB.Find(&mail, int(ID))

	c.JSON(200, mail)
}

func (p *MailController) Store(c *gin.Context) {

	var mail models.Mail

	_ = p.DB.Create(&mail)

	c.JSON(200, mail)
}

func (p *MailController) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	mail := models.Mail{}
	_ = p.DB.Delete(&mail, int(ID))

	c.JSON(200, mail)
}
