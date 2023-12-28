package main

import (
	"net/http"

	database "github.com/Wanna-be-Coder/go-users/db"
	"github.com/Wanna-be-Coder/go-users/users"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	database.InitDB()

	router.HandleFunc("/users", users.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", users.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", users.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", users.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":80", router)
}
