package main

import (
	"bytes"
	"fmt"
	"io"
)

//best practices:
//many small interfaces vs large monos
//do export interfaces for types that will be used by the package
//don't export interfaces for types that will be consumed
//Other languages is the opposite, you make interfaces for things you don't implement yourself
//Design functions and methods to recieve interfaces whenever possible

//MAIN BENEFIT, if you're writing a library, you don't have to THINK of interfaces
//Users can create their own interfaces to your structs
//Developer is no longer responsible for creating the usecases of their own application
//all methods that implement a value type, have to have value recievers
//if I'm implementing an interface with a pointer > then I just have the methods there regardless

func main() {
	writerInterfaceExample()
	incrementorExample()
	interfaceConversion()
	emptyInterface()
	typeSwitch()
}

func typeSwitch() {
	var i interface {} = 0
	switch i.(type) {
		case int:
			fmt.Println("i is int")
		case string:
			fmt.Println("i is string")
		default:
			fmt.Println("i is other")
		
	}
}

func emptyInterface() {
	//empty interface
	//everything can be cast to the empty interface
	//useful if multiple things that are not type compatible
	//Need to then typecase or reflect to figure out what it is
	var myObj interface {} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello YouTube Listeners, this is a test"))
		wc.Close()
	}
	//casting this to figure out what you recieved
	reader, ok := myObj.(io.Reader)
	if ok {
		fmt.Print(reader)
	} else {
		fmt.Println("Conversion failed")
	}
}

func interfaceConversion() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube Listeners, this is a test"))
	wc.Close()

	//casting an interface into into BufferedWriterCloser , NEED TO POINT
	bwc := wc.(*BufferedWriterCloser)
	//can test type conversion like so:
	reader, ok := wc.(io.Reader)
	if ok {
		fmt.Println(reader)
	} else {
		fmt.Println("Conversion failed")
	}
	fmt.Println(bwc)
}

type Closer interface {
	Close() error
}

//you just created an interface that has two interfaces on it
//brilliant
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		if _, err := bwc.buffer.Read(v); err != nil {
			return 0, nil
		}
		if _, err := fmt.Println(string(v)); err != nil {
			return 0, nil
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		if _, err := fmt.Println(string(data)); err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

//------------------------------------

type Incrementor interface {
	Increment() int
}

//Normal type that implements the interface
type IntCounter int

//method that int counter has implemented
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func incrementorExample() {
	//castign integer into an int counter
	myInt := IntCounter(0)
	//this has to be a pointer
	var inc Incrementor = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

//---------------------------------------------------

//interface type
//Naming convention is to name it buy what it does + er
//Multi method interfaces get a bit more confusing
type Writer interface {
	Write([]byte) (int, error)
}

//don't need to explicitly IMPLEMENT the interface
type ConsoleWriter struct {}

//METHOD for ConsoleWriter that implements Writer interface
func (cw ConsoleWriter) Write(data []byte) (int,error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func writerInterfaceExample() {
	fmt.Println("Hello, playground")
	//this is holding the INTERFACE in console writer struct you instantiated
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}