package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	js := fmt.Sprintf(`"status":"available","environment":%q,"version":%q`, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
