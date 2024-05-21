package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Duma-D/simple-http-server-go/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	persons []dto.PersonDTO
}

func NewServer() *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		persons: []dto.PersonDTO{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/create-person-record", s.createPerson()).Methods("POST")
	s.HandleFunc("/list-everybody", s.listPersons()).Methods("GET")
	s.HandleFunc("/delete-person-record", s.removePerson()).Methods("DELETE")
}

func (s *Server) createPerson() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var person dto.PersonDTO
		if err := json.NewDecoder(req.Body).Decode(&person); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("Error decoding body: " + err.Error())
			return
		}
		person.ID = uuid.New()

		s.persons = append(s.persons, person)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listPersons() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.persons); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("Error encoding persons list: " + err.Error())
			return
		}
	}
}

func (s *Server) removePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr, _ := mux.Vars(req)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for i, person := range s.persons {
			if id == person.ID {
				s.persons = append(s.persons[:i], s.persons[i+1:]...)
				break
			}
		}
	}
}
