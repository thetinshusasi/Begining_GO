package main

import "fmt"
// strings create a backing array and object ( ptr to backing array, length of the string)

func main(){
	var arr [5] string
	arr[0]="1" /// pointer of 1 is copied to arr[0] , this is perfomance improvement
	arr[1]="2"
	arr[2]="3"
	arr[3]="4"
	arr[4]="5"
	for i:=0; i< len(arr);i++{
		fmt.Println(arr[i])
	}

	for i, val:=range(arr){
		val = val + "tinshu"
		fmt.Println(i," ==> " ,val)

	}

	for i:=0; i< len(arr);i++{
		fmt.Println(arr[i])
	}


	/// array literal
	arr2 := []int{
		10,
		20,
		30,
	}

	for i, val:=range(arr2){
		fmt.Println(i," ==> " ,val)

	}
}