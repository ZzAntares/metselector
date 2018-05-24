package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/ZzAntares/metselector/internal/models"
)

func (app *App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func (app *App) DatabaseCheckHandler(w http.ResponseWriter, r *http.Request) {
	var q models.Question

	// modify structure to take the mongo ID and remove the published at
	c := app.Database.C("questions")
	oerr := c.Find(bson.M{}).One(&q)

	if oerr != nil {
		log.Fatal(oerr)
	}

	json.NewEncoder(w).Encode(q)
}

func (app *App) QuestionsListHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.AllQuestions(app.Database))
}

func (app *App) SuggestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Call OK!")
}
