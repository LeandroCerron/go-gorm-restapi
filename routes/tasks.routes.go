package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/leandroCerron/go-gorm-restapi/db"
	"github.com/leandroCerron/go-gorm-restapi/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" || idParam == "0" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id field is required and cannot be zero"))
		return
	}

	taskId, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	var task models.Task
	db.DB.First(&task, taskId)
	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	if task.UserId == 0 || task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please complete all fields"))
		return
	}

	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id field is required"))
		return
	}

	taskDeleted := db.DB.Unscoped().Delete(&task, id)
	err := taskDeleted.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
