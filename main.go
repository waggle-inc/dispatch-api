// Package main is the Waggle Inc. dispatch service.
//
// Assigns delivery jobs to available drivers. Part of Waggle Freight.
package main

import "fmt"

// deliveriesPerDriver is the standard route size for a single shift.
const deliveriesPerDriver = 4

// AssignDeliveries returns the number of deliveries that can be
// scheduled for the given number of available drivers.
func AssignDeliveries(drivers int) int {
	if drivers <= 0 {
		return 0
	}
	return drivers * deliveriesPerDriver
}

func main() {
	drivers := 3
	fmt.Println("Waggle dispatch-api starting")
	fmt.Printf("Drivers available: %d\n", drivers)
	fmt.Printf("Deliveries scheduled: %d\n", AssignDeliveries(drivers))
}
