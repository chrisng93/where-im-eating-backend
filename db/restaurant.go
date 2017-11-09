package db

import (
	"gopkg.in/mgo.v2/bson"
)

func (ops Ops) GetRestaurants() ([]Restaurant, error) {
	var results []Restaurant
	err := ops.collection.Find(nil).All(&results)
	return results, err
}

func (ops Ops) GetRestaurant(id string) (Restaurant, error) {
	result := Restaurant{}
	err := ops.collection.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

func (ops Ops) InsertRestaurant(restaurant Restaurant) (*bson.ObjectId, error) {
	err := ops.collection.Insert(&restaurant)
	restaurant.ID = bson.NewObjectId()
	_, err = ops.collection.UpsertId(restaurant.ID, &restaurant)
	if err != nil {
		return nil, err
	}
	return &restaurant.ID, nil
}

func (ops Ops) UpdateRestaurant(id string, restaurant Restaurant) error {
	_, err := ops.collection.UpsertId(bson.ObjectIdHex(id), &restaurant)
	return err
}

func (ops Ops) DeleteRestaurant(id string) error {
	return ops.collection.RemoveId(bson.ObjectIdHex(id))
}
