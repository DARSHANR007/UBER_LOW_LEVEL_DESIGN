package main

type DriverType int

const (
	Normal DriverType = iota
	Premium
)

type Driver struct {
	Name            string
	Age             int
	ID              string
	Phone           string
	Email           string
	Address         string
	License         string
	Vehicle         string
	Status          bool
	Type            DriverType
	Lat             float64
	Long            float64
	currentCustomer *Customer
}

func (d *Driver) CurrentCustomerDetails(c *Customer) {

	d.currentCustomer = c

}
