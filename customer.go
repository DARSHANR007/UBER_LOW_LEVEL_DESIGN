package main

type Choice int

const (
	Normal Choice = iota
	Premium
	Suv
)

type Customer struct {
	Name    string
	Age     int
	Address string
	Phone   string
	Email   string
	ID      string
	Type    Choice
}

func (c *Customer) bookRide(Location string, Type string) (int, float64, float64) {

}
