package main

import (
	"fmt"
)

func ifStatment() {
	//Need curly braces
	if true {
		fmt.Println("Test is true")
	}

	mapExample := map[string]int {
		"Hi": 1,
	}
	//if statement on pop, ok
	if pop, ok := mapExample["Hi"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 30
	if guess < number || guess > number {
		fmt.Println("or operator")
	} 
	if guess < number && guess > number {
		fmt.Println("and operator")
	} 
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("Got it!")
	}
	//flips bool
	fmt.Println(!true)

	//short circuiting
	//when one part of or test returns true it immediately passes
	//same with And test, if one piece fails it exits early

	//if, elif, else 
	if 1 + 1 != 2 {
		fmt.Println("Got it!")
	} else if 1+2 == 2 {
		fmt.Println("Got it!")
	} else {
		fmt.Println("Got it!")
	}
}

func switchStatement() {
	//does requires to have breaks in switch cases
	//if you want your cases to fall through use the keywork fallthrough
	//fallthrough happens even if the other case SHOULDNT execute
	//Break keyword to leave switch early
	
	switch 2 {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("not one or Two")
	}
	//can have multiple values in a case
	switch 2 {
		case 1 , 4:
			fmt.Println("one")
		case 2, 5:
			fmt.Println("two")
		default:
			fmt.Println("not one or Two")
	}
	//can initialize in switch
	switch i := 2+2; i {
	case 1 , 4:
		fmt.Println("one")
	case 2, 5:
		fmt.Println("two")
	default:
		fmt.Println("not one or Two")
	}
	//can switch over types
	var i interface{} = 1
	switch i.(type) {
		case int:
			fmt.Println("one")
		case string:
			fmt.Println("two")
		default:
			fmt.Println("not one or Two")
	}
}

func main() {
	ifStatment()
}