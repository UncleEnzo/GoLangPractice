package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//best practices:
//don't use goroutines in libraries
//know how it's going to end
//check for race conditions at compile time
//use go run -race to whatever you're invoking.  This gives you data races

//syncronizes goRoutines together
var wg = sync.WaitGroup{}
var counter = 0
//anything can read, but only one can write,
//waits till readers are done before writing
var m = sync.RWMutex{}
//same as above but any can read or write
var z = sync.Mutex{}

func main() {
	//by default os threads = to cores on machine
	//can use this method to set the number of threads you want
	runtime.GOMAXPROCS(100)
	basicsGoRoutine()
	waitGroupExample()
	unreliableWaitGroup()
	mutexExample()
}

func mutexExample() {
	//Terrible use in this example because it accomplishes nothing faster
	//but if you use this parralelism with more complicated tasks it'll make them faster
	for i := 0; i < 10; i++ {
		wg.Add(2)
		//lock outside of goroutines, unlock inside of them
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func unreliableWaitGroup(){
	//this makes it unreliable | No sync between the routines
	for i := 0; i < 10; i++ {
		wg.Add(2) // two new go routines are being made right now
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func waitGroupExample(){
	//better solution: wait group
	var msgTwo = "hello"
	//adds extra value to wait group counter
	wg.Add(1)
	go func(msgTwo string) {
		fmt.Println(msgTwo)
		//finish parrallel thread wg
		wg.Done()
	}(msgTwo)
	//waits until all counters are done in wait group
	wg.Wait()
}

func basicsGoRoutine(){
	//makes it a go routine, which parrallelizes it
	//creating threads is expensive in other language
	//in go goroutines are not expensive, they're abstracted
	// so easy create/destroy
	go sayHello()
	//Need to wait on this so here's a horrible action

	//can invoke anon:
	var msg = "hello"
	//closure: anon go routines have access to higher variables > This is a bad idea
	go func() {
		fmt.Println(msg)
	}()
	//Race condition here: 
	msg = "Goodbye" // since the go function executes async from the main function
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}