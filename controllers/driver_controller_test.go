package controllers

import (
	"testing"

	mgo "gopkg.in/mgo.v2"

	"gitlab.com/varver/wmd/config"
	"gitlab.com/varver/wmd/db"
	"gitlab.com/varver/wmd/models"
	"gopkg.in/mgo.v2/bson"
)

func mockCollection() *mgo.Collection {
	config.Setting.EnvMode = "test"
	session := db.Database()
	c := session.DB(config.Setting.TestDatabase).C(config.Setting.DriversLocationCollection)
	return c
}

func Test_ValidDriverID(t *testing.T) {

	if ValidDriverID(2) {
		t.Log("ValidDriverID test passed for id = 2")
	} else {
		t.Error("Error in valid driver id for id = 2")
	}

	if !ValidDriverID(60000) {
		t.Log("ValidDriverID test passed for id = 60,000")

	} else {
		t.Error("Error in valid driver id for id = 60,000")
	}
}

func Test_upsertDriver(t *testing.T) {
	c := mockCollection()
	c.RemoveAll(bson.M{})
	id := int64(42)
	data := models.Driver{}
	data.ID = id
	data.Accuracy = 0.5
	data.Location.Type = "Point"
	data.Location.Coordinates = []float64{77.135431, 28.737728}
	_, err := upsertDriver(c, id, data)
	if err != nil {
		t.Error("Error in inserting driver location : " + err.Error())
	}

	data.Location.Coordinates = []float64{77.135231, 28.737428}
	_, err = upsertDriver(c, id, data)
	if err != nil {
		t.Error("Error in updating driver location : " + err.Error())
	}

}

func Test_getDrivers(t *testing.T) {
	c := mockCollection()
	input := DriversRequestWrapper{
		28.737468,
		77.135241,
		0,
		2000,
	}
	drivers, err := getDrivers(c, input)
	if err != nil {
		t.Error("Error in geting driver's list : " + err.Error())
	}

	if len(drivers) != 1 {
		t.Error("Number of driver records fetched are wrong")
	}
}
