package main
import "fmt"

// every go routine is set of instruction to be excuted  the cpu
// every go routine is given around 2kB memory for its stack , normally in other prog lang its 1 MB per thread
// when stack is full, then new stack is allocated with double the size of the prev stack and prev stack is copied over (doesn't happen often)
// once stack size usage is reduced to 25 % of its  allocation , GC will do another copy operation to remove the extra space


func main(){

	count:=10
	fmt.Println( "count ", count, &count)
	increment(count)
	fmt.Println( "increment ", count, &count)
	increment1(&count)
	fmt.Println( "increment1", count, &count)
	increment(count)
	fmt.Println( "address reused after returning ", count, &count)

	/// escape analysis go determines whether value stay in stack or heap
	var ex example
	println("value before intialization", ex.intVal)

	ex = valueOnStack()
	println("value on stack", ex.intVal)
	ex= *(escapeToStack())
	println("value on heap", ex.intVal)

	///  two go routine cann't share the same stack




}

func increment( inc int){
	inc++;
	fmt.Println( "inc ", inc, &inc)

}

func increment1( inc *int){
	*inc++;

	fmt.Println( "inc ", *inc , &inc)

}

type example struct {
	floatVal float64
	intVal   int
	boolVal  bool
}

func valueOnStack() example{
	ex := example {
		floatVal: 2.2,
		boolVal: true,
		intVal:2,
	}
/// stack are self cleaning i.e this memory will be reused once again by another func call by excuting Go routine
	return ex;
}
func escapeToStack() *example{
	///  this data will be put on heap and "ex" will have  the pointer  to it
	ex := example {
		floatVal: 2.2,
		boolVal: true,
		intVal:4,
	}
/// escape analysis of go push this to heap

	return &ex;
}