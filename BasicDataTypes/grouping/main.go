package main

import "fmt"

/// In Go , we cann't group by

type Animal struct {
	Think string
}

type Cat struct {
	Animal
	Lives int
}

type Dog struct {
	Animal
	Bite bool
}

func main() {

	/// this won't work as subtyping doesn't suport diversity , though this is may go against oops principle
	/// however this is "is a "  relationship or inheritance relationship
	/// we should be propmoting  "has a " relationship" or interface or behaviour grouping
	// 	daisy: Cat{
	// 		Animal: Animal{
	// 			Think: "jmes",
	// 		},
	// 		Lives: 10,
	// 	},
	// 	tony: Dog{
	// 		Animal: Animal{
	// 			Think: "i like u",
	// 		},
	// 		Bite: true,
	// 	},
	// }

	for _, v := range animals {
		fmt.Println(v.Think)
	}

}
