package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/ZzAntares/metselector/internal/models"
)

type App struct {
	Database *mgo.Database
	// TODO Move this to a shared package?
	// TODO Add a reference to the logger?
}

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
	var questions []models.Question
	// TODO Read DB settings from environment
	c := app.Database.C("questions")
	oerr := c.Find(bson.M{}).All(&questions)

	if oerr != nil {
		log.Fatal("Can't read from database, check permissions and resource names")
		return
	}

	json.NewEncoder(w).Encode(questions)
}
