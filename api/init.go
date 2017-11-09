package api

import (
	"net/http"

	"github.com/chrisng93/where-im-eating-backend/db"
	"github.com/gorilla/mux"
)

func Init(dbOps db.Ops) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/restaurant", func(w http.ResponseWriter, r *http.Request) { getRestaurantsHandler(w, r, dbOps) }).Methods("GET")
	r.HandleFunc("/restaurant", func(w http.ResponseWriter, r *http.Request) { addRestaurantHandler(w, r, dbOps) }).Methods("POST")
	r.HandleFunc("/restaurant/{id}", func(w http.ResponseWriter, r *http.Request) { getRestaurantHandler(w, r, dbOps) }).Methods("GET")
	r.HandleFunc("/restaurant/{id}", func(w http.ResponseWriter, r *http.Request) { updateRestaurantHandler(w, r, dbOps) }).Methods("PUT")
	r.HandleFunc("/restaurant/{id}", func(w http.ResponseWriter, r *http.Request) { deleteRestaurantHandler(w, r, dbOps) }).Methods("DELETE")
	return r
}
