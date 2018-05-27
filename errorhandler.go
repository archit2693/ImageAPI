package main

import(
	"encoding/json"
	"net/http"
)


type JSONError struct {
     ErrorMessage string
     Status uint
}

func ReportError(w http.ResponseWriter, error_message string, status uint){
     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
     error := JSONError{ErrorMessage: error_message, Status: status}
     json.NewEncoder(w).Encode(error)
}
