package main

import (
	"fmt"
)

/// basic types in Go
func main() {
	// size of int is based on the pointer size / Word size of the hardware its running on
	var integer int
	fmt.Println(integer)
	var floatVal float64
	fmt.Println(floatVal)
	fmt.Printf("%T \n", integer)

	/// there is no casting in GO there are only conversions
	floatVal = 100.0
	integer = int(floatVal)
	fmt.Println(floatVal)

	var exampleStruct example
	println(exampleStruct.floatVal)

	/// struct literal
	exampleStruct = example{
		boolVal:  true,
		floatVal: 2.1,
		intVal:   2,
	}
	println(exampleStruct.floatVal)
	exampleStruct2 := example{
		boolVal:  true,
		floatVal: 200.45,
		intVal:   2,
	}
	println(exampleStruct2.floatVal)

	/// Anonmyous Struct
	ananomyousStruct := struct {
		intVal int
	}{
		intVal: 30,
	}

	println(ananomyousStruct.intVal)

	/// Implict conversation between named type and annomyous type is allowed,
	/// However not betweeen two named types

	var ex example
	// var ex2 example2
	// println(ex,ex2)
	//ex= ex2

	ex3 := struct {
		/// Always keep the highest data type to lowest , due to the concept of padding in GO
		floatVal float64
		intVal   int
		boolVal  bool
	}{
		floatVal: 100.1,
		intVal:   2,
		boolVal:  true,
	}

	ex = ex3
	println(ex.floatVal)

}

/// Create a custom type in GO
type example struct {
	/// Always keep the highest data type to lowest , due to the concept of padding in GO
	floatVal float64
	intVal   int
	boolVal  bool
}
type example2 struct {
	/// Always keep the highest data type to lowest , due to the concept of padding in GO
	floatVal float64
	intVal   int
	boolVal  bool
}
