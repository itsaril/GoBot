package handlers

import (
	"encoding/json"
	"net/http"

	"gobot/database/repository"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := repository.GetQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}
