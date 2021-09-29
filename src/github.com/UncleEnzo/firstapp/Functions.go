package main

import (
	"fmt"
)

//example of multiple in a row being same type
func sayGreeting(greeting, name string) string {
	return greeting + name;
}

//passing in pointers
func sayGreetingPointers(greeting, name *string) string {
	return *greeting + *name;
}

//variatic function, must be last param
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(result)
}

//return reference as a pointer
func sumReferenceReturn(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

//names return values
func sumReferenceReturnNames(values ...int) (result int) {
	fmt.Println(values)
	// result := 0 dont need to declare
	for _, v := range values {
		result += v
	}
	return //don't need to name the return either
}

//error handling
func divide(a,b float64) (float64, error) {
	if b == 0.0 { 
		return 0.0, fmt.Errorf("Error")
	}
	return a / b, nil
}

//methods > this is a function that's executed in a known type
type greeter struct {
	greeting string
	name string
}

//method | Gets copy of a struct > can do this pointers of course to avoid copying
func (greeterStruct greeter) greet() {
	fmt.Println(greeterStruct.greeting, greeterStruct.name)
}

func main() {
	fmt.Println(sayGreeting("Hello", "stacey"))

	name := "Nick"
	greeting := "Hi"
	fmt.Println(sayGreetingPointers(&greeting, &name))

	//error handling pattern
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//anon function, immediately invoked | no purpose other than isolated scope
	func() {
		fmt.Println("Hello go")
	}()

	//save and use later func
	f := func() {
		fmt.Println("Hello go")
	}
	var x func() = func() {
		fmt.Println("Hello go")
	}
	f()
	x()

	

	g := greeter {
		greeting: "hello",
		name: "Go",
	}
	g.greet()
}