package utils

import (
	"errors"
	"math"
)

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// distance returned is METERS!!!!!!
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

//ValidateLatLong returns if the provided latitude and longitude are valid or not ?
func ValidateLatLong(lat float64, long float64) (err error) {
	if lat > float64(90.0) || lat < float64(-90.0) {
		err = errors.New("Latitude should be between +/-90")
		return
	}
	if long > float64(180) || long < float64(-180) {
		err = errors.New("Longitude should be between +/-180")
	}
	return
}
