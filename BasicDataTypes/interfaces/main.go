package main

import (
	"BeginingGo/BasicDataTypes/embedding"
	"fmt"
)

type Reader interface {
	read(b []byte) (int, error)

	/// this is a way of thinking in go (think in the perceptive of the user),
	// better one is above as the backing bytes array is passed user,  not backing array is created by the read method
	// passed up the call stack
	///read(i int) ( [] byte, error)

	/// interfaces are referrance types and two word Data structures
	/// each represent two diff pointers
	/// 1) points to actual or copy ( based on pass by ref or value)
	/// 2) point to I-Table that hold info and location abt the methods on the current exceuting type and current type
	/// Learn abt the naming convention of interface , such reader interface will do read op
	/// when we are trying to set concrete type to interface type and then trying to call the concrete type reciever methods
	/// via interface variable . Then method call success will depend upon signature of receiver method
	/// ie value receiver methods can be called by both by  pointer ref and value ref but not vice versa
}

type Writer interface {
	write(b []byte) (int, error)
}

/// interface composition
type ReaderWriter interface {
	Reader
	Writer
}

// create concrete type struct when u will be using them to store state or share state
//  try to aggregate based in behaviour
// create subtypes for type promotions
// look at the struct u are creating to decide whether it is aggregation struct or concrete struct or both
type User struct {
	Name string
	Age  int
	Reader
	Writer
}

func (u *User) write(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func (u *User) read(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func main() {
	fmt.Println("interfaces")
	a := embedding.User{
		FirstName: "james",
	}
	fmt.Println(a.FirstName)
}
