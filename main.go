package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/ZzAntares/metselector/controllers"
	"github.com/ZzAntares/metselector/models"
)

func main() {
	// Load settings file
	err := godotenv.Load()
	if err != nil {
		log.Println("*** Unable to load settings file ***")
		log.Fatal(err)
	}

	// Initialize database
	database, err := models.NewDBConnection(os.Getenv("MONGODB_NAME"))
	if err != nil {
		log.Panic(err)
	}
	defer database.Session.Close()

	app := &controllers.App{Database: database}

	// Router initialization
	var router = mux.NewRouter()
	// TODO Move this to a routes file?
	router.HandleFunc("/healthcheck", app.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/questions", app.QuestionsListHandler).Methods("GET")
	router.HandleFunc("/suggest", app.SuggestHandler).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	middlewareCORS := handlers.CORS(originsOk, headersOk, methodsOk)

	severAddress := os.Getenv("APP_BIND_ADDRESS")
	severPort := os.Getenv("APP_BIND_PORT")
	log.Println("Starting application...")
	log.Printf("Listening on http://%s:%s/\n", severAddress, severPort)

	// Start application server
	log.Fatal(http.ListenAndServe(
		severAddress+":"+severPort,
		middlewareCORS(router)))
}
