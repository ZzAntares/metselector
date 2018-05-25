package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZzAntares/metselector/controllers"
	"github.com/ZzAntares/metselector/models"
)

type mockDB struct{}

func (mdb *mockDB) AllQuestions() []*models.Question {
	return make([]*models.Question, 0)
}

// Take this as a sample test
func TestHealthCheckHandler(t *testing.T) {
	rresponse := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

	app := &controllers.App{Database: &mockDB{}}
	http.HandlerFunc(app.HealthCheckHandler).ServeHTTP(rresponse, req)

	var expected string = "\"Still alive!\"\n"
	var obtained string = rresponse.Body.String()

	if expected != obtained {
		t.Errorf(
			"\n...expected = %#v\n...obtained = %#v",
			expected,
			obtained)
	}
}
