package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user_profile/database"
)

type Handler struct {
	DB *sql.DB
}
type data map[string]interface{}

func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		fmt.Fprintf(w, "Welcome!")
		return
	}
}

func (h *Handler) Userprofile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}
	username := r.Header.Get("username")
	if username == "" {
		JSONWriter(w, data{
			"Error": "no username found in header",
		}, http.StatusInternalServerError)
		return
	}

	response, err := database.Getuserdetails(username, h.DB)
	if err != nil {
		JSONWriter(w, data{
			"Error": err,
		}, http.StatusInternalServerError)
		return
	}

	JSONWriter(w, data{"Response": response}, 200)
}
