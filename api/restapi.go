package api

import (
	"fmt"
	"net/http"
)

// CreatePerson /app/person/create
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreatePerson")
}
