package main

import (
	"errors"
	"fmt"
	"reflect"
)

///	 it is prefeable to use Errors values from the errors package
/// however u can always create custom error types in go lang , all you have to do is implement  four of these
/// error interface with concrete error types to decouple the errors from concrete types
/// by making concrete types unexported
/// This is not an hard and fast rule
/// 1) Temperoray
/// 2) Not Found
/// 3) Timeout
/// 4) Not Authorized

/// implement 4 methods

/// Always use error interface for the return error values
/// and don't use any other values set to its zero value , as u will be returning a ptr of Type set to its nill value
/// which is still a value
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
