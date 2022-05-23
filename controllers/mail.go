package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fredmessias43/rest-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MailController struct {
	DB *gorm.DB
}

var mails = []models.Mail{
	{ID: 1, Content: "John", Subject: "Doe"},
	{ID: 2, Content: "Koko", Subject: "Doe"},
}

func (m *MailController) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mails)
}

func (m *MailController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range mails {
		ID, _ := strconv.ParseInt(params["id"], 0, 32)
		if item.ID == int(ID) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Mail{})
}

func (m *MailController) Store(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var mail models.Mail
	_ = json.NewDecoder(r.Body).Decode(&mail)
	ID, _ := strconv.ParseInt(params["id"], 0, 32)
	mail.ID = int(ID)
	mails = append(mails, mail)
	json.NewEncoder(w).Encode(mails)
}

func (m *MailController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range mails {
		ID, _ := strconv.ParseInt(params["id"], 0, 32)
		if item.ID == int(ID) {
			mails = append(mails[:index], mails[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(mails)
	}
}
