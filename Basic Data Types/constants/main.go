package main

import "fmt"
func main(){

	///minimum precision is 256 bits 
	const pie = 22.0/7
	fmt.Println(pie)

	/// precision is lost at  named type constants
	const pie2 float64 = 22.0/7
	fmt.Println(pie2)


	/// Enums i believe i Go iota
	const(
		A1 =20	// 20
		B1 		// 20	
		C1 = iota 	//2
		D1			//3
	)
	fmt.Println(A1,B1,C1, D1)

	const(
		A2 = 1<< iota	// 1
		B2 		// 2
		C2  	//4
		D2			//8
	)
	fmt.Println(A2,B2,C2, D2)

}