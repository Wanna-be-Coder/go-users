package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "github.com/Wanna-be-Coder/go-users/db"
	"github.com/gorilla/mux"
)

type CreateUserParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Insert the new user into the database
	createdUser, error := database.InsertUser(newUser.Name, newUser.Age)
	if error != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the entire created user object in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	user, err := database.GetUserByID(userID)
	if err != nil {
		if err == database.ErrNoUser {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Fatal(err)
		}
		return
	}

	// Convert user to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

type UpdateUserParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreateUserParams

	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Insert the new user into the database
	updatedUser, error := database.UpdateUser(userID, newUser.Name, newUser.Age)
	if error != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the entire created user object in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedUser)

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	err = database.DeleteUserByID(userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	// Convert user to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}
