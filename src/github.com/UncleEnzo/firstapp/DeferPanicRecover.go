package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferExample() {
    //if you defer all of them it'll be backwards: LIFO order, last in first out
	//Reason for lifo is to close resources in opposite order you opened them
	fmt.Println("Start")
	// > executes this before return after the rest of the function if deferred
	defer fmt.Println("Middle") 
	fmt.Println("End")
}

func deferHttpResourcePattern() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	//closes it right next to opening so that you can do all your work 
	//and not forget to close the resource later
	//not a good idea in a loop
	defer res.Body.Close()
	//takes stream and parses out into string of bytes
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func deferReassignedVar() {
	//this actually print start at the end of the function
	//not end because defer captures the value when it's originally called
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func panicExample() {
	//many exceptions that we have in other languages aren't exceptions in go
	//Ex, opening files that don't exist is an error NOT an exception in go
	//panic is CANT CONTINUE TO FUNCTION
	//automatically throws a panic
	a, b := 1,0
	ans := a/b
	//appropriately generates the panic
	fmt.Println(ans)
}

func createdPanic() {
	fmt.Println("Start")
	panic("Something bad happened")
}

func panicPracticalHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	//used to Error out if you open another terminal and run this app
	//while it's already running
	if err := http.ListenAndServe(":8080", nil); err != nil {
		//We are deciding that this will crash the program
		//Errors on their own do not
		panic(err.Error())
	}
}

func panicAfterDefer() {
	fmt.Println("start")
	//defers get called before panics
	//defer takes function CALL/invokation
	defer func() {
		//recover makes functions higher up the call stack still execute after
		//THIS function exits at panic.  like throwing errors I guess in Java
		//BUT they continue to work
		if err := recover(); err != nil {
			log.Println("Eror:", err)
			//this is to repanic the application if it 
			//fails to be handled higher up by recover
			panic(err)
		}
	}()
	panic("something bad happened")
}

func main() {
	deferExample()
	deferHttpResourcePattern()
	deferReassignedVar()
	panicExample()
	createdPanic()
	panicPracticalHttp()
	panicAfterDefer()
}