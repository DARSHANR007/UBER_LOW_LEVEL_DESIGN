package main

type payment struct {
	paymentType   string
	paymentID     string
	amount        int
	status        bool
	paymentDate   string
	paymentTime   string
	paymentMethod string
}

func (p *payment) completePayment() bool {
	return true
}
