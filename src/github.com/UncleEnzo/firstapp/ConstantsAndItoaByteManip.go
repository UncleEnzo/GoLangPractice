package main

import "fmt"

//consts are PascalCase for exporting
//camelCase for package
//typed constants work like immutable variables
//untyped constants work like literals and interoperate when added to other types Ex: const 10 + int32(10) will add together
//iota enumerates but STARTS at 0 > this is why you use the blank in the beginning of KB example so the first expression iota = 1
//Enumerated expressions: iota with constants + arithmetic or bitwaise operations or bitshifting

const (
	_ = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func byteConverter() {
	fileSize := 4000000000.
	//format result 2 decimal places and literal string gb
	fmt.Printf("%.2fGB", fileSize/GB)
}

//Sets boolean flags inside of a single byte
//Iota is a counter for all of the constants in the block
//in this example the counter is used with bit shifting the value one
const (
	isAdmin = 1 << iota //000001
	isHeadquarters //000010
	canSeeFinancials //000100
	canSeeAfrica //etc
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func isAdminExample() {
	//assigns all of these to roles because each one is resolved in the bit
	//100101 > Meaning those 3 are true
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b/n", roles)
	//bitmask use to return roles and admin if true check
	fmt.Printf("Is Admin? %v", isAdmin & roles)
}


func main() {
	isAdminExample()
	byteConverter()
}