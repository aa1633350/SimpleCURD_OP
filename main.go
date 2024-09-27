package main

import (
	"SimpleCURD_OP/goroutines"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/people", createPersonHandler)
	r.Get("/people", getAllPeopleHandler)
	r.Get("/people/{id}", getPersonHandler)
	r.Put("/people/{id}", updatePersonHandler)
	r.Delete("/people/{id}", deletePersonHandler)
	goroutines.DemoGoRoutine()
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {

}

func updatePersonHandler(w http.ResponseWriter, r *http.Request) {

}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {

}

func getAllPeopleHandler(w http.ResponseWriter, r *http.Request) {

}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {

}
