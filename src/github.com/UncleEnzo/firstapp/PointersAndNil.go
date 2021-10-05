package main

import (
	"fmt"
)

func newVars(){
	a := 42
	b := a
	fmt.Println(a , b)
	a = 27
	//b == 42 a == 27
	fmt.Println(a, b)
}

func pointers(){
	a := 42
	//b pointer to interger address of a
	//* here means point to a Address
	var b *int = &a
	fmt.Println(a , b)
	a = 27
	//Sets B to point to A so both change
	//b prints out computer memory location of A
	fmt.Println(a, b)
	//Asterisk here means resolve address to location
	// Dereference when putting asterisk in front
	//both equal 27
	fmt.Println(a, *b)
	//this will change the value of both A and B's pointers
	*b = 14
	fmt.Println(a, *b)
}

func unsafePointerArithmatic() {
	//go doesn't allow for pointer memory
	//If you REALLY NEED pointer arithmatic, use the "unsafe" package
	a := [3]int {1, 2, 3}
	b := &a[0]
	c := &a[1] // - 4 > not allower but lets you jump between addresses insanely fast
	fmt.Println(a, b, c)
}

func structAddressing() {
	type myStruct struct {
		foo int
	}

	//this says ms is holding the address of an object that has 
	//a field with a value 42 in it
	var ms *myStruct
	ms = &myStruct { foo: 42 }
	fmt.Println(ms)
}

func nilExample() {
	//pointer you don't initialize are initilized as nil
	
}

func main() {
	newVars()
	pointers()
	unsafePointerArithmatic()
}