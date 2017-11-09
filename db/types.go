package db

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const RestaurantCollection = "Restaurant"

type Restaurant struct {
	ID              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name            string        `json:"name"`
	Region          string        `json:"region"`
	City            string        `json:"city"`
	Cuisine         string        `json:"cuisine"`
	Summary         string        `json:"summary"`
	OrderSuggestion string        `json:"order_summary"`
	YelpURL         string        `json:"yelp_url"`
}

type Ops struct {
	session    *mgo.Session
	collection *mgo.Collection
	name       string
}
