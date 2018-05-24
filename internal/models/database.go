package models

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

type Datastore interface {
	AllQuestions() []*Question
}

type DB struct {
	*mgo.Database
}

// TODO Move to a more general file/package?
func NewDBConnection(databaseName string) (*DB, error) {
	// TODO Take into consideration user, password
	//      See: https://godoc.org/github.com/globalsign/mgo#Dial
	session, err := mgo.Dial(
		os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT"))

	if err != nil {
		log.Println("*** Unable to connect to the database ***")
		return nil, err
	}

	if err = session.Ping(); err != nil {
		log.Println("*** Connected to database, but can't ping it ***")
		return nil, err
	}

	return &DB{session.DB(databaseName)}, nil
}
