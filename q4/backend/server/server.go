package server

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	Db     *sql.DB
	Router *mux.Router
	Port   int
}

func (s *Server) Init() {
	s.Router.HandleFunc("/users", s.GetUsers).Methods("GET")
	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users/{id}", s.GetUser).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
}

func (s *Server) ServeHTTP() {
	http.ListenAndServe(":"+strconv.Itoa(s.Port), &CORSRouterDecorator{s.Router})
}
