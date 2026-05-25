// package main

// import (
// 	method "practice/methods"
// )

// func main() {
// 	// fmt.Println("..........Let Get Things Started 😋..........")
// 	// var age int = 10
// 	// var text string = "Shola is "
// 	// method.ModifiedPointer(&age, &text)
// 	// fmt.Print("\n-------------------------------------------------\n\n")
// 	// // method.WorkingSlices()
// 	// fmt.Println(method.Solution("world"))

// 	// arr1 := []bool{
//     //   true,  true,  true,  false,
//     //   true,  true,  true,  true ,
//     //   true,  false, true,  false,
//     //   true,  false, false, true ,
//     //   true,  true,  true,  true ,
//     //   false, false, true,  true,
//     // }
// 	// fmt.Println(method.CountSheeps(arr1))

// 	// char := "bkko"
// 	// fmt.Println(method.RemoveChar(char))

// 	method.PrintDetails()

// }

package main

import "fmt"

// the main idea of the project is to be able to do something else when there is a panicmessage from our program, instead of displaying the panic to the console.

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something went wrong") //this causes a panic automatically
	}

	fmt.Println("All is well")
}

func toRecover() {
	defer func() {
		if r := recover(); r != nil { //recover function catches the panic, then do something else.
			fmt.Println("There was a panic: ")
		}
	}()

	mightPanic(true)
}

func main() {
	toRecover()
}
