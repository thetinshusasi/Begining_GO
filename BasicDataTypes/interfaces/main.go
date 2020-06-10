package main

import (
	"BeginingGo/BasicDataTypes/embedding"
	"fmt"
)

type reader interface {
	read(b []byte) (int, error)

	/// this is a way of thinking in go (think in the perceptive of the user),
	// better one is above as the backing bytes array is passed user,  not backing array is created by the read method
	// passed up the call stack
	///read(i int) ( [] byte, error)

	/// interfaces are referrance types and two word Data structures
	/// each represent two diff pointers
	/// 1) points to actual or copy ( based on pass by ref or value)
	/// 2) point to I-Table that hold info and location abt the methods on the current exceuting type and current type
}

func main() {
	fmt.Println("interfaces")
	a := embedding.User{
		FirstName: "james",
	}
	fmt.Println(a.FirstName)
}
