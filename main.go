package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/controllers"
	"github.com/fredmessias43/rest-api/models"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Person{}, &models.Mail{})

	person := controllers.PersonController{DB: db}
	router.GET("/people", person.Index)
	router.GET("/people/:id", person.Show)
	router.POST("/people", person.Store)
	router.DELETE("/people/:id", person.Delete)

	mail := controllers.MailController{DB: db}
	router.GET("/mails", mail.Index)
	router.GET("/mails/:id", mail.Show)
	router.POST("/mails", mail.Store)
	router.DELETE("/mails/:id", mail.Delete)

	router.Run()
}
