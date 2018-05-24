package controllers

import "github.com/globalsign/mgo"

type App struct {
	Database *mgo.Database
	// TODO Move this to a shared package?
	// TODO Add a reference to the logger?
}
