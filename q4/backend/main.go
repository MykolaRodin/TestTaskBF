package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	Routers()
}

func Routers() {
	err := connectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	http.ListenAndServe(":9080", &CORSRouterDecorator{router})
}

/***************************************************/

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id,first_name,last_name,email from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]

	sqlStatement := `INSERT INTO users(first_name,last_name,email) VALUES($1,$2,$3)`
	_, err = db.Exec(sqlStatement, first_name, last_name, email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New user was created")
}

// Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, first_name,last_name,email from users WHERE id = $1", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]

	sqlStatement := `UPDATE users SET first_name = $1, last_name= $2, email=$3 WHERE id = $4`
	_, err = db.Exec(sqlStatement, first_name, last_name, email, params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}

/***************************************************/

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

//Db configuration
var db *sql.DB

func connectDB() error {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:secretpassword@127.0.0.1:5432/gousers?sslmode=disable")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
