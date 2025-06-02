package main

import (
	"fmt"
)

type paymenter interface {
	payment(amount float32)  // payment method that takes an amount as a float32
}

type payment struct {
	gateway paymenter // paymenter interface type
}

/* func (p *payment) makePayment(amount float32) {
	razorpayPayment := razorpay{}
	razorpayPayment.payment(amount)
} */

type razorpay struct {

}

func (r *razorpay) payment(amount float32) {
	fmt.Println("Payment made using Razorpay: ", amount)
}

type stripe struct {
	// Stripe-specific fields
}
func (s *stripe) payment(amount float32) {
	fmt.Println("Payment made using Stripe: ", amount)
}
func main() {
	fmt.Println("Interfaces in Go")
	// Interfaces are used to define a contract that structs can implement.
	// They allow for polymorphism, where different types can be treated as the same type if they implement the same interface.

	newPayment := razorpay{}
	newPayment.payment(100.50)

	paymentStrip := stripe{}
	paymentStrip.payment(200.75) // Using Stripe payment method
	fmt.Println("Payment made successfully")
}