package main

import (
	"fmt"
	"sync"
	"time"
)

//most languages have concurrency and parralellism as
//3rd party stuff channels pass data between go routines

var wg = sync.WaitGroup{}

func main() {
	basicChannel()
	deadLockChannel()
	twoWayChannel()
	directionLockedChannel()
	bufferedChannels()
	bufferedChannelsRecieverClose()
	selectStatements() 
}

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
//struct with no field require 0 memory.
//signal only channel
//do this over boolean, saves memory
var doneCh = make(chan struct {})

func selectStatements() {
	//monitors log channel for log entries coming from application
	go logger()
	//need to close this example 1
	// defer func() {
	// 	close(logCh)
	// }()
	logCh <- logEntry{ time.Now(), logInfo, "App is starting" }
	logCh <- logEntry{ time.Now(), logInfo, "App is shutting down" }
	//pass message into done channel when you want it to shut down
	//this syntax in struct is saying define empty struct, fill it with nothing
	doneCh <- struct{}{}
}

func logger() {
	for {
		//select lets you either expect the log
		//or a shut down signal from the done channel
		select {
		case entry := <-logCh:
				fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <- doneCh:
			break
		}
		//can have default but then it will be NON blocking
	}
}

func bufferedChannelsRecieverClose() {
	//second param gives buffer for storing info in channel if push more than pull
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		//checks channel to make sure it's closed
		//recieves data from a channel but not in a loop
		for {
			if i, ok := <- ch; ok {
			fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		//send 2
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func bufferedChannels(){
	//second param gives buffer for storing info in channel if push more than pull
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		//recieves from buffered channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		//send 2
		ch <- 42
		ch <- 27
		//signals that nothing else is coming in from the channel
		//if you dont' have, for loop will get error
		//can't pass into closed channel
		//can't detect that channel is closed
		//use recover to check if pushing to closed channels
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func directionLockedChannel() {
	ch := make(chan int)
	//direction lock channel by forcing dir with parameter
	wg.Add(2)
	//recieve only
	go func(ch <-chan int) {
		i := 42
		fmt.Println(i)
		//so cant do this
		// ch <- 27
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		//and can't do this here
		// fmt.Println(<-ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func twoWayChannel() {
	ch := make(chan int)

	//both goroutines are readers and writers in this case.
	//bad practice though, should have dedicated read/writer
	wg.Add(2)
	go func() {
		i := 42
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

func deadLockChannel() {
	ch := make(chan int)

	//this causes deadlock because channel is sending first and 
	//closes while others are still pulling from it
	//sending
	wg.Add(1)
	go func() {
		i := 42
		//going into channel, value 42
		ch <- i
		i = 27
		wg.Done()
	}()

	for j := 0; j < 5; j++ {
		wg.Add(1)
		//recieving data
		go func() {
			//This line pauses goroutine until the it can get something
			//from the channel.  If no open pushing channel, this causes deadlock
			i := <- ch
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func basicChannel() {
	//channels must be make with make
	//channels are strongly typed
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		//recieving data
		go func() {
			//recieving from channel into variable
			i := <- ch
			fmt.Println(i)
			wg.Done()
		}()

		//sending
		go func() {
			i := 42
			//going into channel, value 42
			ch <- i
			i = 27
			wg.Done()
		}()
	}
	wg.Wait()
}