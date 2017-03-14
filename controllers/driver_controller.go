package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gitlab.com/varver/wmd/config"
	"gitlab.com/varver/wmd/logger"
	"gitlab.com/varver/wmd/models"
	"gitlab.com/varver/wmd/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	DriverController struct {
		session *mgo.Session
	}
)

func NewDriverController(s *mgo.Session) *DriverController {
	return &DriverController{s}
}

type DriversRequestWrapper struct {
	Latitude  float64
	Longitude float64
	Limit     int64
	Radius    int64
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type DriverResponse struct {
	Latitude  float64
	Longitude float64
	ID        int64
	Distance  float64
}

func ValidDriverID(id int64) bool {
	// check range of driver id
	if id < config.Setting.DriverIdStart || id > config.Setting.DriverIdEnd {
		return false
	}
	return true
}

//upsertDriver will either insert or update the current location of a driver in database
func upsertDriver(c *mgo.Collection, id int64, obj models.Driver) (interface{}, error) {

	// check latitude and longitude
	points := obj.Location.Coordinates
	err := utils.ValidateLatLong(points[1], points[0])
	if err != nil {
		return nil, err
	}

	// finally insert/update data
	obj.Time = time.Now()
	change, err := c.UpsertId(id, obj)
	return change, err
}

func getDrivers(c *mgo.Collection, input DriversRequestWrapper) ([]models.Driver, error) {
	var results []models.Driver
	err := utils.ValidateLatLong(input.Latitude, input.Longitude)
	if err != nil {
		logger.Infof("%v", err)
		return results, err
	}

	if input.Radius == 0 {
		input.Radius = config.Setting.DefaltRadius
	}
	if input.Limit == 0 {
		input.Limit = config.Setting.DefaultResultsLimit
	}

	// fetch all the drivers within provided radius and limit
	err = c.Find(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{input.Longitude, input.Latitude},
				},
				"$maxDistance": input.Radius,
			},
		},
	}).Limit(int(input.Limit)).All(&results)

	if err != nil {
		return results, err
	}

	return results, nil
}

func makeDriversResponse(params DriversRequestWrapper, drivers []models.Driver) []DriverResponse {
	driversResponse := []DriverResponse{}
	for _, val := range drivers {
		temp := DriverResponse{}
		temp.ID = val.ID
		logger.Infof("%v\n", val.Location.Coordinates)
		long := val.Location.Coordinates[0]
		lat := val.Location.Coordinates[1]
		temp.Latitude = lat
		temp.Longitude = long
		temp.Distance = utils.Distance(params.Latitude, params.Longitude, lat, long)
		driversResponse = append(driversResponse, temp)
	}
	return driversResponse
}

func (dc *DriverController) FetchDrivers(r render.Render, req *http.Request) {
	input := DriversRequestWrapper{}
	v := req.URL.Query()
	allErrors := ErrorResponse{}
	for key, value := range v {
		if len(value) > 0 {
			key = strings.ToLower(key)
			if key == "latitude" {
				lat, err1 := strconv.ParseFloat(value[0], 64)
				if err1 != nil {
					allErrors.Errors = append(allErrors.Errors, "Parameter latitude is required")
					continue
				}
				input.Latitude = lat
			} else if key == "longitude" {
				long, err1 := strconv.ParseFloat(value[0], 64)
				if err1 != nil {
					allErrors.Errors = append(allErrors.Errors, "Parameter latitude is required")
					continue
				}
				input.Longitude = long
			} else if key == "limit" {
				limit, err1 := strconv.ParseInt(value[0], 10, 64)
				if err1 != nil {
					allErrors.Errors = append(allErrors.Errors, "Parameter limit's value should be integer")
					continue
				}
				input.Limit = limit
			} else if key == "radius" {
				radius, err1 := strconv.ParseInt(value[0], 10, 64)
				if err1 != nil {
					allErrors.Errors = append(allErrors.Errors, "Parameter radius value should be integer")
					continue
				}
				if radius < 1 {
					allErrors.Errors = append(allErrors.Errors, "Radius should be greater than zero (0)")
					continue
				}

				input.Radius = radius
			}
		}
	}

	// if any error till here, throw it
	if len(allErrors.Errors) > 0 {
		logger.Errf("%v", allErrors)
		r.JSON(400, allErrors)
		return
	}

	c := dc.session.DB(config.Setting.Database).C(config.Setting.DriversLocationCollection)
	drivers, err2 := getDrivers(c, input)
	if err2 != nil {
		allErrors.Errors = append(allErrors.Errors, err2.Error())
		logger.Crit("Something went wrong in getting driver's data : " + err2.Error())
		r.JSON(400, allErrors)
		return
	}

	finalResponse := makeDriversResponse(input, drivers)
	r.JSON(200, finalResponse)
	return
}

func (dc *DriverController) SaveDriverLocation(r render.Render, req *http.Request, params martini.Params) {

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil || !ValidDriverID(id) {
		logger.Err("invalid user id, must be integer")
		r.JSON(404, make(map[string]string))
		return
	}

	latitude, longitude, accuracy := req.FormValue("latitude"), req.FormValue("longitude"), req.FormValue("accuracy")

	allErrors := ErrorResponse{}

	// basic checks part 1
	if len(latitude) < 1 {
		allErrors.Errors = append(allErrors.Errors, "Parameter latitude is required")
	}
	if len(longitude) < 1 {
		allErrors.Errors = append(allErrors.Errors, "Parameter longitude is required")
	}
	if len(accuracy) < 1 {
		allErrors.Errors = append(allErrors.Errors, "Parameter accuracy is required")
	}
	// if any error till here, throw it
	if len(allErrors.Errors) > 0 {
		r.JSON(422, allErrors)
		return
	}

	// basic checks part 2
	lat, err1 := strconv.ParseFloat(latitude, 64)
	if err1 != nil {
		allErrors.Errors = append(allErrors.Errors, "Parameter latitude must be int/float")
	}
	long, err1 := strconv.ParseFloat(longitude, 64)
	if err1 != nil {
		allErrors.Errors = append(allErrors.Errors, "Parameter longitude must be int/float")
	}

	acq, err1 := strconv.ParseFloat(accuracy, 64)
	if err1 != nil {
		logger.Info(err1.Error())
		allErrors.Errors = append(allErrors.Errors, "Parameter accuracy must be int/float")
	} else {
		if acq < 0 || acq > 1.0 {
			allErrors.Errors = append(allErrors.Errors, "Accuracy should be between 0.0 - 1.0 inclusive")
		}
	}
	// if any error till here, throw it
	if len(allErrors.Errors) > 0 {
		r.JSON(422, allErrors)
		return
	}

	logger.Infof("ID => %v , lat => %v , long => %v , acc => %v", id, lat, long, acq)

	c := dc.session.DB(config.Setting.Database).C(config.Setting.DriversLocationCollection)
	driverObj := models.Driver{}
	driverObj.Accuracy = acq
	driverObj.ID = id
	driverObj.Location.Type = "Point"
	driverObj.Location.Coordinates = []float64{long, lat}

	// insert or update driver's data
	_, err = upsertDriver(c, id, driverObj)
	if err != nil {
		msg := fmt.Sprintf("Problem in updating driver information, %s", err.Error())
		allErrors.Errors = append(allErrors.Errors, msg)
		r.JSON(422, allErrors)
		return
	}

	// else return success with empty body
	r.JSON(200, make(map[string]string))
	return
}
