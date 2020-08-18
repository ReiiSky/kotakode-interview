package models

import "github.com/Kamva/mgm/v3"

// Activity models
type Activity struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `bson:"title" json:"title"`
	Priority         int    `bson:"priority"  json:"priority"`
	Detail           string `bson:"detail"  json:"detail"`
	Group            string `bson:"group" json:"group"`
}

// GetCollection return the collection of current Activity model
func (act *Activity) GetCollection() *mgm.Collection {
	return mgm.Coll(act)
}
