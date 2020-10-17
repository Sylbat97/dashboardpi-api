package handler

import (
	"log"
	"net/http"
)

// RootHandler interface use as a wrapper around the handler functions.
type RootHandler func(http.ResponseWriter, *http.Request) error

// rootHandler implements http.Handler interface.
func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	error, ok := err.(BaseHttpError) // Check if it is a http error.
	if !ok {
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := error.ResponseBody() // Try to get response body of error.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := error.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
