package db

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

func Init(username string, password string, host string, name string) Ops {
	var uri string
	if username != "" && password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s", username, password, host, name)
	} else {
		uri = fmt.Sprintf("mongodb://%s/%s", host, name)
	}
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Error connecting to Mongo: %v", err)
	}
	c := session.DB(name).C(RestaurantCollection)
	return Ops{session: session, collection: c, name: name}
}
