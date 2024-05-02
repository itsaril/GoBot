package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"gobot/api/handlers"
	"gobot/api/middleware"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login).Methods("POST")

	questionRouter := r.PathPrefix("/questions").Subrouter()
	questionRouter.Use(middleware.AuthMiddleware)
	questionRouter.HandleFunc("/", handlers.GetQuestions).Methods("GET")

	http.Handle("/", r)
}
