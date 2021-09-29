package main

import (
	"fmt"
)

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i + 1, j + 2 {
		fmt.Println(i,j)
	}

	//scope I outside of the loop, can you do this in C#?
	i := 0
	for ; i < 5; {
		fmt.Println(i)
		i++
	}

	//nested loops, if you break from inside it goes one loop up, so to break out of 
	//an entire loop you use a label
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(j)
			if j == 3 {
				break Loop
			}
		}
		fmt.Println(i)
	}

	//while loop
	y := 0
	for {
		fmt.Println("hello")
		if y == 2 {
			continue
		}
		if y == 5 {
			break
		}
	}
}

func forRange() {
	s := []int{1,2,3}
	//key and value of collection | index and value in this case
	//works for slices, array, string, and maps
	for k, v := range s {
		fmt.Println(k,v)
	}

	//key only loop
	for k := range s {
		fmt.Println(k)
	}

	//value only loop | since otherwise the variable would just be the key
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	forLoop()
}