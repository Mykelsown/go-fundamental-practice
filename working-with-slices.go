package practice

// import "fmt"

// func WorkingSlices() {
// 	listOfNum := []int{1, 2, 3, 4 ,5, 6 ,7, 8, 9}
// 	fmt.Printf("%d\n", listOfNum)

// 	s1 := listOfNum[4:]
// 	fmt.Printf("s1 = %d, type = %T", s1, s1)
// }

func Solution(word string) string {
  newWord := ""
  for l := len(word)-1; l >= 0; l-- {
    newWord = newWord + string(word[l])
  } 
  
  return newWord
}

func CountSheeps(numbers []bool) int {
  ans := 0
  for _, v := range numbers {
		if bool(v) == true {
			ans++
    }
  }
  
  return ans
}

func RemoveChar(word string) string {
  newWord := ""
  for i, l := range word{
    if i == 0 || i == len(word)-1 {
      continue
    }
    newWord += string(l)
  }
  return newWord
}

