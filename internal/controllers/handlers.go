package controllers

import (
	"encoding/json"
	"net/http"
)

func (app *App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func (app *App) QuestionsListHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(app.Database.AllQuestions())
}

func (app *App) SuggestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Call OK!")
}
