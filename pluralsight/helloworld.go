package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//Go will not allow a non-declaration statement at the package level.
//Global variable declaration
var (
	/*
	  name string //name of subscriber
	  course string //Name of the current course
	  module float64 //Current place int he course
	*/
	//Method 2
	//name, course, module = "Brian", "Deep dive", 3.4
	//Method 3
	name   string  = "Brian" //name of subscriber
	course string  = "Brian" //Name of the current course
	module float64 = 3.5     //Current place int he course

	asdf string = "asdf"
)

func main() {
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("Name is set to", name, "and is of type: ", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type: ", reflect.TypeOf(module))

	//Declare and initialize variable syntax. Shorthand only works in functions
	a := 10.000000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and b is of type: ", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC has value", c, "and c is of type: ", reflect.TypeOf(c))

	//unusedVariableHere := "lalalala"

	ptr := &a

	fmt.Println("\nMemory address of pointer is", ptr, "and value is: ", *ptr)

	changeCourse(course)
}

func changeCourse(course string) string {
	return ""
}
