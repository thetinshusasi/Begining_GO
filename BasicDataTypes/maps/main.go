package main

import "fmt"

func main(){
	maps:= make(map[string]string)
	maps["t"]="tinshu"

	for k ,v := range(maps){
		fmt.Println(k, " ===> ", v)
	}

	mapsLiteral:= map[string]string{
		"s":"sweetys",
		"l":"load",
		}

	for k ,v := range(mapsLiteral){
		fmt.Println(k, " ===> ", v)
	}

	delete(mapsLiteral,"l")
	fmt.Println("deleting operation")

	for k ,v := range(mapsLiteral){
		fmt.Println(k, " ===> ", v)
	}

	fmt.Println("maps reterival  operation")

	u, isExist := mapsLiteral["s"]
	fmt.Println(u, " ===> ", isExist)
	u, isExist = mapsLiteral["l"]
	fmt.Println(u, " ===> ", isExist)



}