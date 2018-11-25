package api

import (
	"encoding/json"
	"net/http"
)

// Hello response structure
type Hello struct {
	Message string
}

// HelloHandleFunc to be used as http.HandleFunc for Hello API
// w: to give response to client
// r: Request we received
func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	m := Hello{"Welcome to Cloud Native Go."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Hello cloud native go")
	w.Write(b)
}
