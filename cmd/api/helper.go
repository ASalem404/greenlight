package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

// We can then use the ByName() method to get the value of the "id" parameter from
// the slice. In our project all movies will have a unique positive integer ID, but
// the value returned by ByName() is always a string. So we try to convert it to a
// base 10 integer (with a bit size of 64). If the parameter couldn't be converted,
// or is less than 1, we know the ID is invalid so we use the http.NotFound()
// function to return a 404 Not Found response

// Retrieve the "id" URL parameter from the current request context, then convert it to
// an integer and return it. If the operation isn't successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {

	// js, err := json.Marshal(data)

	// Use the json.MarshalIndent() function instead json.Marshal() so that whitespace is added to the encoded
	// JSON. Here we use no line prefix ("") and tab indents ("\t") for each element.

	/***
	While using json.MarshalIndent() is positive from a readability and user-experience point
	of view, it unfortunately doesn’t come for free. As well as the fact that the responses are
	now slightly larger in terms of total bytes, the extra work that Go does to add the
	whitespace has a notable impact on performance.
	***/

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
