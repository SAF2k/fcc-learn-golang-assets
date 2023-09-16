package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 0
	var costPerSMS float64 = 0.000000
	var hasPermission bool = false
	var username string = ""

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
