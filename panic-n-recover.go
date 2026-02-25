package practice

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something went wrong")
	}

	fmt.Println("All is well")
}

func ToRecover() {
	defer func (){
		if r := recover(); r != nil {
			fmt.Println("There was a panic: ")
		}
	}()

	mightPanic(true)
}

