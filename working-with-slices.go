package practice

import (
	"fmt"
)

func WorkingSlices() {
	listOfNum := []int{1, 2, 3, 4 ,5, 6 ,7, 8, 9}
	fmt.Printf("%d\n", listOfNum)

	s1 := listOfNum[4:]
	fmt.Printf("s1 = %d, type = %T", s1, s1)
}