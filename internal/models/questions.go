package models

import "github.com/globalsign/mgo/bson"

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
