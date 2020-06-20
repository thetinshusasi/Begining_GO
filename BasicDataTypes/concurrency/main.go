package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("concurrency ")

	can := make(chan string)

	go func() {
		can <- "james"

	}()

	go func() {
		fmt.Println("I am orphan")
		time.Sleep(2 * time.Second)
		can <- "picka"
		time.Sleep(2 * time.Second)
		can <- "picka"

		fmt.Println("I am orphan 2")

	}()
	// fmt.Println(<-can)
	// can = nil

	fmt.Println("done")
	fmt.Println(<-can)
	fmt.Println(<-can)
	fmt.Println(<-can)

}

/// there are four major parts where go scheduler can make a decison
// // 1) go call
// // 2) sys calls
// // 3) chanel op
// // 4) garbage collector

// as rule of thumb, optimize for correctness , if its then only
// optimize to use more cores.
// try not to make more a balance for go routines
