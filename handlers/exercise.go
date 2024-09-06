package handlers

import (
	"fmt"
	"net/http"
)

func POST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Using Posty Posty")
}

func GET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Using Getty Getty")
}


func exerciseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		POST(w, r)
	case http.MethodGet:
		GET(w,r)
	default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}