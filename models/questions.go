package models

import (
	"log"
	"reflect"

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

func (db *DB) AllQuestions() []*Question {
	var questions []Question

	c := db.C("questions")
	err := c.Find(bson.M{}).All(&questions)

	if err != nil {
		log.Fatal("Can't read from database, check permissions and resource names")
	}

	// TODO When used multiple-times move it to a utils/shared package
	// Used to convert from "*[]Question" to "[]*Question"
	pointersOf := func(v interface{}) interface{} {
		in := reflect.ValueOf(v)
		out := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(in.Type().Elem())), in.Len(), in.Len())
		for i := 0; i < in.Len(); i++ {
			out.Index(i).Set(in.Index(i).Addr())
		}
		return out.Interface()
	}

	return pointersOf(questions).([]*Question)
}
