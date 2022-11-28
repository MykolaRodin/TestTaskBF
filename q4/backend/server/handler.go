package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"backend/model"
)

// Get all users
func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := s.Db.Query("SELECT id,first_name,last_name,email from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	var users []model.User
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Create user
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]

	sqlStatement := `INSERT INTO users(first_name,last_name,email) VALUES($1,$2,$3)`
	_, err = s.Db.Exec(sqlStatement, first_name, last_name, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "New user was created")
}

// Get user by ID
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := s.Db.Query("SELECT id, first_name,last_name,email from users WHERE id = $1", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	var user model.User
	for result.Next() {
		err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(user)
}

// Update user
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]

	sqlStatement := `UPDATE users SET first_name = $1, last_name= $2, email=$3 WHERE id = $4`
	_, err = s.Db.Exec(sqlStatement, first_name, last_name, email, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

// Delete User
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := s.Db.Exec(sqlStatement, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}
