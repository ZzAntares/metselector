package models

import (
	"log"

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

func (db *DB) AllQuestions() []Question {
	var questions []Question

	c := db.C("questions")
	err := c.Find(bson.M{}).All(&questions)

	if err != nil {
		log.Fatal("Can't read from database, check permissions and resource names")
	}

	return questions
}
