package users

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User created")

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User found")

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User updated")

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User deleted")

}
