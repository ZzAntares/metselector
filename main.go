package main

import (
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/ZzAntares/metselector/internal/controllers"
)

func main() {
	// Load settings file
	err := godotenv.Load()
	if err != nil {
		log.Println("*** Unable to load settings file ***")
		log.Fatal(err)
	}

	// Initialize database
	session, err := mgo.Dial(
		os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT"))
	if err != nil {
		log.Println("*** Unable to connect to the database ***")
		log.Fatal(err)
	}

	app := &controllers.App{Database: session.DB(os.Getenv("MONGODB_NAME"))}

	// Router initialization
	var router = mux.NewRouter()
	// TODO Move this to a routes file?
	router.HandleFunc("/healthcheck", app.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/questions", app.QuestionsListHandler).Methods("GET")
	router.HandleFunc("/dbcheck", app.DatabaseCheckHandler).Methods("GET")

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
