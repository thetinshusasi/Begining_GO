package main

import (
	"errors"
	"fmt"
	"reflect"
)

///	 it is prefeable to use Errors values from the errors package
/// however u can always create custom error types in go lang , all you have to do is implement error interface
///


/// implement 4 methods
type UnmarshingError struct {
	Type reflect.Type
}

func (e *UnmarshingError) Error() string {
	return "Type : " + e.Type.String()
}
func main() {
	var (
		ErrorBadReq = errors.New("Bad Request")
	)
	fmt.Println(ErrorBadReq)

}
