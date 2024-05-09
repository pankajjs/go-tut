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

	// composite data types
	// array, slice, map
	// declaration
	// array => [4]int
	// slice => []int
	// map => map[string] int

	// array are typed by size, which is fixed at compile time
	
	// all these are equivalent
	// var a [3]int
	// var b [3]int{0, 0, 0}
	// var c [...]{0, 0, 0} // sized by initializer

	// var d [3]int
	// d = b //elements are copied, not reference
	
	// var m [...]int{1, 2, 3, 4}

	// c = m // types mismatched
	
	// Arrays are passed by value, thus elements are copied

	// slices, slices works the same way as string does
	// It has a descriptor, lenth and capacity
	// var a []int   // nil
	// var b []int{1, 2} // initialized
	
	// a = append(a, 1)    append 1 to a 
	// b = append(b, 3)	  []int{1, 2, 3}

	// a = b copied b descriptor in a
	
	// d := make([]int, 5) // []int{0, 0, 0, 0, 0}
	// e := a  //same storage
	// e[0] == a[0] // true

	// slices are passed by reference
	// slices are indexed like [8: 11] 8 is included and 11 is not
	
	// maps
	// maps are dictionaries: indexed by key, returning a value
	// read from a nil map, but inserting will panic
	// var m map[string]int // nil, can read but cant insert
	
	// p := make(map[string]int) // non nill but empty
	// a:= p["the"] // return 0 if key is not there
	// b := m["the"] // return 0
	// m["and"] = 1  // panic because nil map
	// m = p // valid => copies p descriptor to m
	// m["and"]++ //OK
	// c := p["and"] // returns 1

	// maps are passed by reference
	// the type used for key nust have == and != defined

	//maps cant be compared to one another
	// maps can be compared only to nil as a special case
	

	// maps have a special two result lookup function
	// the second variable tell if the key was there
	// p := map[string]int{} // non-nil but empty
	// a := p["the"] // return 0
	
	// b, ok := p["and"] // returns 0, false
	
	// p["the"]++
	// c, ok := p["the"]   return 1, true
	

	// control structure
	// if then else
	// All if then statements require braces

	// if a == b {
	// 	// does this
	// }else{
	// 	// does this
	// }

	// another short declaration
	// if err := doSomething(); err != nil {
	// 	return err
	// }

	// go has only for loop
	// for loops
	// three parts, all optional (initialize, check, increment)
	// for i:= 0; i < 10; i++ {

	// }
	// one variable is index for array and key for map
	// for i := range myArray {

	// }
	// // two values index, and value for array and key, value for map
	// for i, v := range myArray {

	// }

	// i, j := 0, 3
	// for {
	// 	if i > j {
	// 		break
	// 	}

	// 	i++
	// }
	

	// only want values, not index
	// for _, v : range myArray {

	// }

	// switch
	// switch a:= f.Get(); a {
		// case 0, 1, 2:

		// case 3, 4, 5, 6:

		// default :

	// }

	// break is not required

	// package
	// everything in go lives inside a package.

	//It's either at package scope or function scope
	// can declare anything at package scope
	// Every name that's capitalized in a package is exported(public) 

	// imports
	// each source file in your package must import what it needs
	// cant have cyclic dependency
	// 

	// declaration
	// Constant declaration with const
	// Type declaration with type
	// variable declaration with var
	// short, initialized varibale decalration of any type := (inside a function)
	// function declaration with func
	// 

	// var a int // default 0
	// var b int = 1


	// var c = 1 // int
	// var d = 1.0 // float
		
	// declaration block
	// var (
	// 	x, y int
	// 	z float64
	// 	s string
	// )


	// short declaration
	// :=
	// cant be used outsidd of a function
	//  must be used in a control statement
	// must declare at least one new varibale
	// wont reuse an existing declaration from an outer scope

	// err := doSomething();

	// err := doSomething() // wrong, redeclare

	// x, err := getSomevalue();   // ok, assign in error


	// structural typing
	// same type if it has the same structure or behavior
	// arrays of the same size and base type
	// slices with the same base type
	// maps of the same sequence of field types
	// structs with the same sequence of field names/types
	// functions with the same parameter & return types

	// a := [...]int{1, 2, 3}
	// b := [3]int{}

	// a = b // ok
	// 

	// Named typing
	// only same type if it has same declared type name
	//  type x int // user defined types

	// func main() {
		// var a x  // x is defined type; base int
		// b := 12 // int
		// a = b // type mismatch
		// a = 12  // ok, untyped literal
		// a = x(b) //ok, type conversion
	// }

}
