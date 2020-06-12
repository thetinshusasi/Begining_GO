package main

import (
	"fmt"
)

type Reader interface {
	read(b []byte) (int, error)
}

type Writer interface {
	write(b []byte) (int, error)
}
type User struct {
	Name string
	Age  int
}

type Admin struct {
	User
}

func (u User) write(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func (u User) read(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func (u Admin) write(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func (u Admin) read(b []byte) (int, error) {

	var er error = nil
	val := 100
	return val, er

}

func main() {
	u := User{
		Age:  10,
		Name: "tinshu",
	}

	a := Admin{
		User: u,
	}

	var wr Writer = u

	/// this is a syntax to do type assertions
	fmt.Println(wr.(User))
	// fmt.Println(wr.(Admin)) /// this fails panic: interface conversion: main.Writer is main.User, not main.Admin
	b, ok := wr.(Admin) /// code doesn't panic here
	///
	fmt.Println(ok, b)
	fmt.Println(a)

	// way to stop th prg
	// log.Fatal()

	//os.Exit(-1)

}
