package main

import (
	"fmt"
	"practice"
)

func main() {
	fmt.Println("..........Let Get Things Started 😋..........")
	var age int = 10
	var text string = "Shola is "
	practice.ModifiedPointer(&age, &text)
	fmt.Print("\n-------------------------------------------------\n\n")
	practice.WorkingSlices()
}