package api

import (
	"net/http"
	"github.com/Duma-D/simple-http-server-go/dto"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	persons []dto.PersonDTO
}

func NewServer() *Server {
	server := &Server{
		Router:  mux.NewRouter(),
		persons: []dto.PersonDTO{},
	}
	server.routes()
	return server
}

func (server *Server) routes() {
	server.HandleFunc("/create-person-record", server.createPersonPost()).Methods("POST")
	server.HandleFunc("/list-everybody", server.listPersonsGet()).Methods("GET")
	server.HandleFunc("/delete-person-record/{id}", server.removePersonDelete()).Methods("DELETE")
}

func (server *Server) createPersonPost() http.HandlerFunc {
	return func(writter http.ResponseWriter, req *http.Request) {
		createPerson(writter, req, server)
	}
}

func (server *Server) listPersonsGet() http.HandlerFunc {
	return func(writter http.ResponseWriter, req *http.Request) {
		listPersons(writter, server)
	}
}

func (server *Server) removePersonDelete() http.HandlerFunc {
	return func(writter http.ResponseWriter, req *http.Request) {
		removePerson(writter, req, server)
	}
}
