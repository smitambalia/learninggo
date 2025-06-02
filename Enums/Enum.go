package main 

import (

	"fmt"
)

type orderStatus string
const (
	Pending orderStatus = "pending" // 0
	Shipped    = "shipped"                 // 1
	Delivered  = "delivered"              // 2
	Cancelled  = "cancelled"              // 3
)
func changeOrderStatus(status orderStatus){
	fmt.Println("Order status : ", status)
}

func main( ) {
	fmt.Println("Enums in Go")
	changeOrderStatus(Shipped) // Change the order status to Shipped
}