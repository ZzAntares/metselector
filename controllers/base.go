package controllers

import "github.com/ZzAntares/metselector/models"

type App struct {
	Database models.Datastore
	// TODO Move this to a shared package?
	// TODO Add a reference to the logger?
}
