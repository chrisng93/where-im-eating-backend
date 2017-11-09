package api

import (
	"encoding/json"
	"net/http"

	"github.com/chrisng93/where-im-eating-backend/db"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func getRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func getRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurant, err := dbOps.GetRestaurant(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stringifed, err := json.Marshal(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(stringifed))
}

func addRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	restaurant := db.Restaurant{}
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var id *bson.ObjectId
	id, err = dbOps.InsertRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var stringified []byte
	responseBody := map[string]string{"id": id.Hex()}
	stringified, err = json.Marshal(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(stringified))
}

func updateRestaurantHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteRestaurantHandler(w http.ResponseWriter, r *http.Request) {

}
