package api

import (
	"github.com/chrisng93/where-im-eating-backend/db"
	"github.com/gorilla/mux"
)

var dbOps db.Ops

func Init(db db.Ops) *mux.Router {
	dbOps = db
	r := mux.NewRouter()
	r.HandleFunc("/restaurant", getRestaurantsHandler).Methods("GET")
	r.HandleFunc("/restaurant", addRestaurantHandler).Methods("POST")
	r.HandleFunc("/restaurant/{id}", getRestaurantHandler).Methods("GET")
	r.HandleFunc("/restaurant/{id}", updateRestaurantHandler).Methods("PUT")
	r.HandleFunc("/restaurant/{id}", deleteRestaurantHandler).Methods("DELETE")
	return r
}
