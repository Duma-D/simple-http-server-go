package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
	"github.com/Duma-D/simple-http-server-go/dto"
)


func createPerson(w http.ResponseWriter, req *http.Request, s *Server) {
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

func listPersons(w http.ResponseWriter, s *Server) {
	w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.persons); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("Error encoding persons list: " + err.Error())
			return
		}
}

func removePerson(writter http.ResponseWriter, req *http.Request, server *Server) {
	idStr, _ := mux.Vars(req)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}
	for i, person := range server.persons {
		if id == person.ID {
			server.persons = append(server.persons[:i], server.persons[i+1:]...)
			break
		}
	}
}
