package models

import "time"

type Driver struct {
	ID       int64     `json:"id" bson:"_id"`
	Accuracy float64   `form:"accuracy" json:"accuracy"`
	Time     time.Time `json:"time" bson:"time"`
	Location GeoJson   `bson:"location" json:"location"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
