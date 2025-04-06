package main

import (
	"math/rand"
)

type Choice int

type Customer struct {
	Name            string
	Age             int
	Address         string
	Phone           string
	Email           string
	ID              int
	Lat             float64
	Long            float64
	Destination     string
	currentLocation string
	currentRide     Ride
	currentDriver   Driver
	status          bool
}

var customerList map[int]*Customer

func createCustomer(Name string, Age int, Address string, Phone string,
	Email string, ID int, Lat float64, Long float64, currentLocation string, Destination string) *Customer {

	newCustomer := &Customer{
		Name:            Name,
		Age:             Age,
		Address:         Address,
		Phone:           Phone,
		Email:           Email,
		ID:              ID, // Convert int ID to string
		Lat:             Lat,
		Long:            Long,
		currentLocation: currentLocation,
		Destination:     Destination,
	}
	customerList[newCustomer.ID] = newCustomer

	return newCustomer
}

func (c *Customer) bookRide(Destination string, currentLocation string) int {
	nearestDriver := bookNearestDriver(c, availableDrivers)

	c.currentDriver = nearestDriver
	nearestDriver.CurrentCustomerDetails(c)
	nearestDriver.Status = true
	c.status = true

	return rand.Intn(9000) + 1000
}
