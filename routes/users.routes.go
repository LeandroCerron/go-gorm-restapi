package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandroCerron/go-gorm-restapi/db"
	"github.com/leandroCerron/go-gorm-restapi/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	userCreated := db.DB.Create(&user)
	err := userCreated.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.Unscoped().Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
