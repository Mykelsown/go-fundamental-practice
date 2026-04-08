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
	// practice.WorkingSlices()
	fmt.Println(practice.Solution("world"))

	arr1 := []bool{
      true,  true,  true,  false,
      true,  true,  true,  true ,
      true,  false, true,  false,
      true,  false, false, true ,
      true,  true,  true,  true ,
      false, false, true,  true,
    }
	fmt.Println(practice.CountSheeps(arr1))

	char := "bkko"
	fmt.Println(practice.RemoveChar(char))

	practice.ToRecover()
}