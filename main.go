package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/leandroCerron/go-gorm-restapi/db"
	"github.com/leandroCerron/go-gorm-restapi/models"
	"github.com/leandroCerron/go-gorm-restapi/routes"
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/task", routes.GetTask).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTasks).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
