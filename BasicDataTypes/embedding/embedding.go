package embedding

import "fmt"

type User struct {
	FirstName string
}

type Admin struct {
	User
	lastName string
}

func (u User) displayFirstName() {
	fmt.Println(u.FirstName)
}
func main() {
	ad := Admin{
		User: User{
			FirstName: "tinshu",
		},
		lastName: "james",
	}
	/// this is embedding  inner type is promoted to use by the outer type
	/// if both inner type and outer type implements the same interface then outer type implementation
	/// takes precedence
	ad.displayFirstName()
}
