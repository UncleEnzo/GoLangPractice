//imports the main package to create an entrypoint to the application
package main

//imports the fmt library to print stuff
import (
	"fmt"
)

//fmt.PrintF formats whatever is inside > the percent sign triggers format
//%v stands for var or const
//%b is byte
//%T is type

const stringTest string = "Hello"

func main() {
	var s string = "this is a string"
	//places the var into the string
	fmt.Printf("%v", s)
	//places the const into the string
	fmt.Printf("%v", stringTest)
	//places the const and it's type into a string separated by ,
	fmt.Printf("%v, %T", stringTest, stringTest)
}