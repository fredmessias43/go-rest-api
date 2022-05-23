package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/fredmessias43/rest-api/models"
)

type PersonController struct {
	DB *gorm.DB
}

func (p *PersonController) Index(w http.ResponseWriter, r *http.Request) {
	people := []models.Person{}
	p.DB.Find(&people)
	println(people)
	json.NewEncoder(w).Encode(people)
}

func (p *PersonController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.ParseInt(params["id"], 0, 32)
	person := models.Person{}
	p.DB.Find(&person, int(ID))

	json.NewEncoder(w).Encode(person)
}

func (p *PersonController) Store(w http.ResponseWriter, r *http.Request) {
	person := models.Person{ID: 3, Firstname: "fred", Lastname: "messias"}

	_ = p.DB.Create(&person)

	json.NewEncoder(w).Encode(person)
}

func (p *PersonController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.ParseInt(params["id"], 0, 32)
	person := models.Person{}
	p.DB.Find(&person, int(ID))

	json.NewEncoder(w).Encode(person)
}
