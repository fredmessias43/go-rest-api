package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/fredmessias43/rest-api/controllers"
	"github.com/fredmessias43/rest-api/models"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	ginLog, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(ginLog)

	router := gin.Default()

	dbLog, _ := os.Create("./logs/db.log")
	out := io.MultiWriter(dbLog)
	newLogger := logger.New(
		log.New(out, "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("./database/test.sqlite"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Person{}, &models.Mail{})

	person := controllers.PersonController{DB: db}
	router.GET("/people", person.Index)
	router.GET("/people/:id", person.Show)
	router.POST("/people", person.Store)
	router.PUT("/people/:id", person.Update)
	router.DELETE("/people/:id", person.Delete)

	mail := controllers.MailController{DB: db}
	router.GET("/mails", mail.Index)
	router.GET("/mails/:id", mail.Show)
	router.POST("/mails", mail.Store)
	router.DELETE("/mails/:id", mail.Delete)

	router.Run()
}
