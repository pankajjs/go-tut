package main

// main package

// import packages
import (
	"fmt"
	"os"
	"strings"
	// "os"
)

func calculateAvg(){
	var sum float64
	var n int

	for {
		var val float64
		_, err := fmt.Fscan(os.Stdin, &val)
		
		if err != nil{
			break
		}

		sum += val
		n++
	}

	fmt.Println("The average is ", sum / float64(n))
}


// entry point of go program
func main(){

	// Single line comment
	/*
	multi line
	comment
	*/
	// print on cmd
	// fmt.Printf("Hello, %s\n", os.Args[1]);

	// data types
	// a := 2 // declare and initialze a integer 
	// a is memory adderess in the machine  unlike interpreted language

	// int is the default type for integers in GO.
	// two types of integer signed and unsigned
	// signed => int64 (default), int32, int16, int8
	// unsigned => uint64(default), uint32, uint16, uint8

	// floating number
	// float32, float64 (default)

	// complex number
	// complex64, complex128

	// simple declaration of varibale
	//Anywhere
	// 	var a int

	// 	// grouping
	// 	var (
	// 		b = 2
	// 		f = 2.01
	// 	)
	// // only inside function
	// 	c := 2 // short declaration 

	// exercise
	// a := 2
	// b := 3.1
	// %T for type and %v for value

	// fmt.Printf("a: %T %v\n", a, a)
	// fmt.Printf("b: %T %v\n", b, b)

	// // align output
	// fmt.Printf("a: %8T %v\n", a, a)
	// fmt.Printf("b: %8T %v\n", b, b)

	// // instead of writing a and b twice, it takes the value and type from
	// // the first argument of printf
	// fmt.Printf("a: %8T %[1]v\n", a)
	// fmt.Printf("b: %8T %[1]v\n", b)
	
	// not allowed, go is very strict
	// a = b

	// explicitly conversion, typecast
	// a = int(b)
	// b = float64(a);
	// fmt.Printf("a: %8T %[1]v\n", a)
	// fmt.Printf("b: %8T %[1]v\n", b)
	
	// bool
	// bool has two values false and true
	// these values are not convertible to integers.

	//error type - special type with one function, Error()
	// it is either nil or non-nil
	// if nil it has error message
	// else not

	//pointers
	// physically address
	// a pointer may be nil or non-nil
	

	// Go initialzes all varibales to "zero" by default if you don't
	// * All numbers get 0 (float 0.0, complex 0i)
	// * bool gets false
	// * string gets "" (empty string)
	// Everything else gets nil: 
		// - pointers
		// - slices
		// - maps
		// - channels
		// - functions
		// - interfaces
	// for aggregate types, all members get their "zero" values

	// constants
	// only numbers, strings and booleans can be constants (immutable)

	// constant can be a literal or a compile time function of a constant 
	

	// data types exercise
	// calculateAvg()


	// strings

	// rune  - sequence of 32bit integer for characters
	// byte - sequence of 8bit unsigned integer for characters
	// string - an immutable sequence of "characters"
	//  physically a sequence of bytes (utf-8)
	// logically a sequence of(unicode) runes

	// Runes(characters) are enclosed in single quotes: 'a'

	// s := "èlite" // (logically len of string-> no of bytes required to represent in unicode characters)
	// fmt.Printf("%T %[1]v %d\n", s, len(s));
	// fmt.Printf("%T %[1]v\n", []rune(s));
	// fmt.Printf("%T %[1]v %d\n", []byte(s), len([]byte(s)));

	// s := "Hello, World"
	// how string represents in memory?
	// Here s is a descriptor, which has a pointer to the string in memmory and the length of string
	// s -> [data, len 12] // data is a pointer which points actual string in memory

	// hello := s[:5]
	// hello points to same characters in memory
	// hello -> [data, len 5]

	// strings are passed by reference, thus they aren;t copied.


	// builtin func
	// s := "a string"
	// x := len(s)
	// strings.Contains(s, 'g')
	// strings.HasPrefix(s, "a")
	// strings.Index(s, "string")


	// s = strings.ToUpper(s) // first copy original string and the change to upper case and returns "A STRING"


	// go garbage collects the previous s string

	




}
