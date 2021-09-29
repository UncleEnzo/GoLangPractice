package main

import (
	"fmt"
	"reflect" //needed for tags
)

func maps(){
	statePopulation := make(map[string]int);
	//key value map
	//Keys have to be tested for equality > slices maps and other functions CANNOT
	//Arrays CAN
	//Return order is not garaunteed
	statePopulation = map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11614373,
	}
	fmt.Println(statePopulation);
	fmt.Println(statePopulation["Ohio"]);
	statePopulation["Georgia"] = 10310371;
	//deleted keys return 0 if you call them again
	delete(statePopulation, "Georgia");
	//try get
	pop, ok := statePopulation["Ohio"]
	fmt.Println(pop, ok)
	fmt.Println(ok)
	fmt.Println(len(statePopulation));
	//This uses pointers so will ahve side effects on the original
	sp := statePopulation;
	delete(sp,"Ohio")
}

type Doctor struct {
	number int
	actorName string
	companions []string
}

func structs() {
	//gathers information related to one concept
	aDoctor := Doctor {
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	//positional syntax, no need to write field name, but don't use, harder to read
	bDoctor := Doctor {
		3,
		"Jon Pertwee",
		[]string {
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(bDoctor)
	//anonymous struct
	cDoctor := struct{name string}{name:"John Pertwee"}
	fmt.Println(cDoctor)
	//independent data sets, creat copies if you make new ones unless you use a pointer
}

//embedding example:
type Animal struct {
	Name string
	origin string
}

//takes animal and is placed inside it
type Bird struct {
	Animal
	SpeekKPH float32
	CanFly bool
}

func structEmbedding(){
	b := Bird{}
	b.Name = "Emu"
	b.origin = "Austraila"
	b.SpeekKPH = 48
	b.CanFly = false
	fmt.Println(b)
	//another way of declaring it
	c := Bird {
		Animal: Animal{ Name: "bird", origin: "AU"},
		SpeekKPH: 1,
		CanFly: false,
	}
	fmt.Println(c)
}

//tagging example
type AnimalTag struct {
	Name string `required max: "100"`
	Origin string
}

func main() {
	maps()
	structs()
	structEmbedding()

	//tagging
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}