package db

import (
	"gopkg.in/mgo.v2/bson"
)

func (ops Ops) InsertRestaurant(restaurant Restaurant) (*bson.ObjectId, error) {
	err := ops.collection.Insert(&restaurant)
	restaurant.ID = bson.NewObjectId()
	_, err = ops.collection.UpsertId(restaurant.ID, &restaurant)
	if err != nil {
		return nil, err
	}
	return &restaurant.ID, nil
}

func (ops Ops) GetRestaurant(id string) (Restaurant, error) {
	result := Restaurant{}
	err := ops.collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return result, err
}
