package main

import "math"

type Ride struct {
	ID                    string
	Driver                *Driver // Changed from string to Driver
	Customer              *Customer
	Status                bool
	Location              string
	Date                  string
	Price                 int
	Payment_Confiramation bool
	Payment_Type          string
	Payment               *payment
}

var availableDrivers map[*Driver]string // Changed from availableRides to availableDrivers

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

func bookNearestDriver(customer *Customer, availableDrivers map[Driver]string) Driver {
	var nearestDriver Driver
	minDistance := math.MaxFloat64

	for driver := range availableDrivers {
		dist := haversine(customer.Lat, customer.Long, driver.Lat, driver.Long)

		if dist < minDistance {
			minDistance = dist
			nearestDriver = driver
		}
	}
	delete(availableDrivers, nearestDriver)

	return nearestDriver
}

func (r *Ride) rideScenario() {

	if r.Customer.Lat == r.Driver.Lat && r.Customer.Long == r.Driver.Long {
		if r.Payment.completePayment() {
			r.Status = true
			r.Payment_Confiramation = true
			availableDrivers[r.Driver] = r.Driver.currentCustomer.Destination
		}

	}
}
