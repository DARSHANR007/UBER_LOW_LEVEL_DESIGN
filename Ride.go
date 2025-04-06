package main

import "math"

type Ride struct {
	ID                    string
	Driver                string
	Customer              string
	Status                bool
	Location              string
	Date                  string
	Price                 int
	Payment_Confiramation bool
	Payment_Type          string
}

var availableRides map[Ride]string

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth's radius in km

	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func bookNearestDriver(customer Customer, availableRides map[Ride]string) Driver {
	var nearestDriver Driver
	minDistance := math.MaxFloat64 // very big number to start with

	for ride := range availableRides {
		driver := ride.Driver
		dist := haversine(customer.Latitude, customer.Longitude, driver.Latitude, driver.Longitude)

		if dist < minDistance {
			minDistance = dist
			nearestDriver = driver
		}
	}

	return nearestDriver
}
