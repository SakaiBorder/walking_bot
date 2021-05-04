package service

import (
	"math"
	"math/rand"
)

type Destination struct {
	Latitude  float64
	Longitude float64
}

func CalcDistance(latitude float64, longitude float64, distance int32) Destination {

	const POLE_RADIUS = 6356752.314245
	latDegree := (360 * 1000) / (2 * math.Pi * POLE_RADIUS) * float64(distance)

	const EQUATOR_RADIUS = 6378137
	longDegree := (360 * 1000) / (2 * math.Pi * (EQUATOR_RADIUS * math.Cos(latitude*math.Pi/180.0))) * float64(distance)

	return Destination{
		Latitude:  latitude + (latDegree * rand.Float64()),
		Longitude: longitude + (longDegree * rand.Float64()),
	}
}
