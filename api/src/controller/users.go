package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users..."))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user by id..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
