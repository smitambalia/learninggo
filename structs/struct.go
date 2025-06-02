package main

import (
	"fmt"
	"time"
)

type order struct {
	orderId int
	amount float32
	status string
	createdAt time.Time // Time when the order was created
	updatedAt time.Time // Time when the order was last updated
	users map[string]string
	
}

type userOrder struct {
	userId   int
	userOrder order // Embedding the order struct
}

func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	fmt.Println("Structs in Go")
	// Structs are used to create complex data types that group together related data.
	// They are similar to classes in other programming languages.

	order1 := order{
		orderId : 1,
		amount : 100.50,
		status : "Pending",
		createdAt : time.Now().UTC(),
		updatedAt : time.Now().UTC(),
		users: map[string]string{
			"firstName": "Smit",
			"lastName": "Patel",
		},
	}
	order1.changeStatus("Shipped") // Change the status of the order
	fmt.Println("Order ID:", order1.orderId)
	fmt.Println("Order Amount:", order1.amount)
	fmt.Println("Order Status:", order1.status)
	fmt.Println("Order Created At:", order1.createdAt)
	fmt.Println("Order Updated At:", order1.updatedAt)
	fmt.Println("Order Users:", order1.users)

	var myOrder order
	myOrder.orderId = 2
	myOrder.amount = 200.75
	myOrder.status = "Completed"	
	myOrder.createdAt = time.Now().UTC()
	myOrder.updatedAt = time.Now().UTC()
	fmt.Println("My Order :", order1)
	fmt.Println("My Order :", myOrder)

	userOrder1 := userOrder{
		userId: 1,
		userOrder: myOrder, // Embedding the order struct
	}

	fmt.Println("User Order ID:", userOrder1.userOrder.orderId)
	fmt.Println("User Order Amount:", userOrder1.userOrder.amount)
	fmt.Println("User Order Status:", userOrder1.userOrder.status)
	fmt.Println("User Order Created At:", userOrder1.userOrder.createdAt)
	fmt.Println("User Order Updated At:", userOrder1.userOrder.updatedAt)
	fmt.Println("User Order Users:", userOrder1.userOrder.users)
	fmt.Println("User ID:", userOrder1.userId)
}