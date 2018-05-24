package models

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Evaluation struct {
	Methodology int `json:"methodology"`
	Evaluation  int `json:"evaluation"`
}

type Question struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Question    string        `json:"question"`
	Criteria    string        `json:"criteria"`
	Evaluations []Evaluation  `json:"evaluations"`
}

func NewDBConnection(databaseName string) (*mgo.Database, error) {
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

	return session.DB(databaseName), nil
}

func AllQuestions(db *mgo.Database) []Question {
	var questions []Question

	c := db.C("questions")
	err := c.Find(bson.M{}).All(&questions)

	if err != nil {
		log.Fatal("Can't read from database, check permissions and resource names")
	}

	return questions
}
