package main

import (
	"fmt"
)

func arrayBasics(){
	//array syntax, can only hold one type of data
	//very fast because contiguous in memory
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v", grades)

	//can replace size with spread operator if you are supplying literals
	gradesLiteral := [...]int{97, 85, 93}
	fmt.Printf("Grades Literal: %v", gradesLiteral)

	//zeroed out array
	var students [3]string
	fmt.Printf("students: %v", students)

	//getting/assigning sytax
	students[0] = "Lisa"
	fmt.Printf("index 0: %v", students)

	//array length
	fmt.Printf("num of students: %v/n", len(students))

	//array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}
	fmt.Println(identityMatrix)	

	//arrays are values in GO, so copying array creates two copies of the data inside
	a := [...]int{1,2,3}
	b:= a
	b[1] = 5
	//these two are 2 different arrays with differnt data inside them, you're not mutating
	fmt.Println(a)
	fmt.Println(b)

	//Pointers
	c := [...]int{1,2,3}
	d:= &c // > this & operator shows that you're pointing to the same data
	d[1] = 5
	//these two point to the same array so it gets mutated
	fmt.Println(c)
	fmt.Println(d)
}

func sliceBasics() {
	// > not specifying size or spread operator in array creates a SLICE
	a:= []int{1,2,3} 
	b := a // >Naturally reference types so no need to use & to point
	b[1] = 5
	fmt.Println(a)
	// fmt.Println(b)
	//has length function as well
	fmt.Printf("Length: %v\n", len(a))
	//has Cap function to see capacity of slice since it resizes
	fmt.Printf("cap: %v\n", cap(a))
}

func sliceCuts() {
	//slice cuts: These are Copies
	//also work with arrays
	a := []int{1,2,3,4,5,6,7,8,9,10}
	b := a[:] //slice of all elements
	c := a[3:] //slice from 4th element to end | inclusive, exclusive values
	d := a[:6] //slice first 6 elements
	e := a[3:6] //slice 4th, 5th, 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func makeFunc() {
	//creates zeroed out array, length 3, capacity 100
	a := make([]int, 3, 100) 
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))
	//reassign to new array since it couldn't fit in the old one
	//puts in value 1, capacity will be 2 (Double length on init through append)
	//used so you can start at certain value and copy without copying the underlying array around
	//Warning: can exceed array size and trigger copy operation, which it does here
	a = append(a,1)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))
	//append is a variatic function, so can add as many params as you want
	//capacity is 8 here, even though it's 5 length because Go doubles the size of 
	//arrays as it reslices
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))
	//slice concatination
	a = append(a, []int{2,3,4,5}...)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))
}

func sliceTrimming(){
	a := []int{1,2,3,4,5}
	//slice from beginning, start at index one
	b := a[1:]
	//slice one off the end
	c := a[:len(a)-1]
	//slice in middle
	d := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//WARNING: These are all the same array, need to loop to make copies
}

func main() {
	arrayBasics()
	sliceBasics()
	sliceCuts()
	sliceTrimming()
}