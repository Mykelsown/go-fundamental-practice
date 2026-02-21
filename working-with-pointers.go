package practice

import (
	"fmt"
	"strconv"
)

func ModifiedPointer(digit *int, text *string) {
	fmt.Println("Yeah Pointers it is 👉🏡...........")

	if digit == nil && text == nil {
		fmt.Println("The Pointers were 'nil'")
		return
	}

	*digit += *digit
	*text += strconv.Itoa(*digit)
	fmt.Printf("The value of the modified digit is %v \n	The value of the modified text is  '%v' \n ", *digit, *text)
}
