package main

import "fmt"

func main() {
	/// make is an inbuilt func ,with following args type* , size* and capacity
	/// reference types  map, func, slice , interface values , channels

	slices := make([]string, 5)
	slices[0] = "hi"
	slices[1] = "hi 1"
	slices[2] = "hi 2"
	slices[3] = "hi 3"
	slices[4] = "hi 4"
	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices))
	}
	/// appending to slice
	slices = append(slices, "hi 5")
	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices))
	}
	/// concat two slices

	slices2 := []string{"hi 6", "hi 7"}

	slices = append(slices, slices2...)
	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices))
	}
	fmt.Println("capacity is increased  for slices and copy operation is performed")
	slices = append(slices, "hi 8")

	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices))
	}

	slices3 := slices[1:5]
	slices3[0] = " james "
	fmt.Println("new slices3  ------------------")
	for i, v := range slices3 {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices3))
	}

	fmt.Println("old slices ------------------")

	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices))
	}

	slices4 := slices[2:6:6]

	for i, v := range slices4 {
		slices4[i] = "pikachu"
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices4), &slices4[i], &slices[i])
	}

	fmt.Println("old slices -------------------")

	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices), &slices[i])
	}
	/// slices are ref types so they will be modified here
	fmt.Println("updating slices -------------------")

	updateSlice(slices)
	for i, v := range slices {
		fmt.Println(i, " ==> ", v, " capacity ", cap(slices), &slices[i])
	}

	name := "tinshu"
	updateString(name)
	println(name)

	fmt.Println(slices)

}

func updateSlice(slice []string) {
	slice[0] = "potter"
}
func updateString(name string) {
	name = "lle"
}
