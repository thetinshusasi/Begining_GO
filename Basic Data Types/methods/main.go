package main

import "fmt"

type user struct{
	id string
	name string
	age int

}

func main(){

	u :=user{
		age:12,
		id:"1",
		name:"tinshu",
	}
	fmt.Println(u.name)
	u.displayUpdatedUser("james")
	fmt.Println(u.name)
	u.displayUpdatedUser1("potter")
	fmt.Println(u.name)
	v:=&u
	v.displayUpdatedUser1("ron") // Go internally fix this to (*v).displayUpdatedUser1("rom")
	fmt.Println(v.name)
	fmt.Println(u.name)


}

func (u user) displayUpdatedUser( name string){
	u.name=name;
	println(u.name)
}


func (u *user) displayUpdatedUser1( name string){
	u.name=name;
	println(u.name)
}