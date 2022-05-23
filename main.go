package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/controllers"
	"github.com/fredmessias43/rest-api/models"
)

func main() {
	router := mux.NewRouter()

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Person{}, &models.Mail{})

	person := controllers.PersonController{
		DB: db,
	}
	router.HandleFunc("/people", person.Index).Methods("GET")
	router.HandleFunc("/people/{id}", person.Show).Methods("GET")
	router.HandleFunc("/people/{id}", person.Store).Methods("POST")
	router.HandleFunc("/people/{id}", person.Delete).Methods("DELETE")

	mail := controllers.MailController{}
	router.HandleFunc("/mails", mail.Index).Methods("GET")
	router.HandleFunc("/mails/{id}", mail.Show).Methods("GET")
	router.HandleFunc("/mails/{id}", mail.Store).Methods("POST")
	router.HandleFunc("/mails/{id}", mail.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
