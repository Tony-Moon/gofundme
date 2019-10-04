package main

import (
	"github.com/190930-UTA-CW-Go/gofundme/fund"
	"fmt"
	"flag"
)

func main() {
	inintbalance := flag.Float64("balance", 100, "default balance")
	initname := flag.String("name", "Tony", "default name")
	flag.Parse()

	fmt.Println(*initname)
	fund := fund.NewFund(*inintbalance)

	// Read the comment at 22 first
	lock := make(chan bool)

	// This go routine created a seporate thread. By default the main is running on a single thread.
	go func() {
		fund.Withdraw(50)
		lock <- true		// Unlocks the channel, does not have to be true, as long as its any message
	} () // () immediately calls the go function
	
	// fmt.Println("Hit ENTER to continue")
	// fmt.Scanln()

	// We have created a "race condition" where the program is unstable. 
	// Some times the go routine will beat the Println, but not always.
	// To fix this we will create a channel up top
	<- lock
	fund.Withdraw(45)

	fmt.Println(fund.Balance())
}