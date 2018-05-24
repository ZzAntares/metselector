package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ZzAntares/metselector/internal/controllers"
)

func main() {
	// Initialize database
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("*** Unable to connect to the database ***")
		log.Fatal(err)
	}

	app := &controllers.App{Database: session.DB("selector_metodologias")}

	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", app.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/questions", app.QuestionsListHandler).Methods("GET")
	router.HandleFunc("/dbcheck", app.DatabaseCheckHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	middlewareCORS := handlers.CORS(originsOk, headersOk, methodsOk)

	log.Println("Starting application...")
	log.Println("Listening on http://localhost:8000/")
	// TODO Read bind address and port from environment
	log.Fatal(http.ListenAndServe(":8000", middlewareCORS(router)))
}
