package main

import "fmt"

// func main() {
// 	str := "my name is aditya"
// 	count := 0
// 	ForInt := make([]int, len(str))
// 	ForStr := make([]string, len(str))
// 	fmt.Println(len(ForInt))
// 	fmt.Println(len(ForStr))
// 	counted := make([]int, len(str))
// 	for i, v := range str {
// 		count++
// 		ForStr[i] = string(v)
// 		// from := count*i
// 		// to := count
// 		// fmt.Println(i)|
// 		if ForStr[i] == " " {
// 			counter := 0
// 			fmt.Println("This is The Count: ", count)
// 			if count == 0 {
// 				continue

// 			}
// 			counted[counter] = count
// 			count = 0
// 		}
// 		// ForInt[i] = int(v)
// 		ForStr[i] = string(v)

// 		// fmt.Println(count)
// 		// if  == emptyStr {

// 		// }

// 	}
// 	fmt.Println(ForInt)
// 	fmt.Println(ForStr)
// 	fmt.Println(counted)
// }

// func main() {
// 	str := "my name is aditya "

// 	currentWord := ""
// 	longestWord := ""
// 	maxLen := 0

// 	for _, ch := range str {
// 		// ch := str[i]

// 		if ch != ' ' {
// 			currentWord += string(ch)

// 		} else {
// 			if len(currentWord) > maxLen {
// 				maxLen = len(currentWord)
// 				longestWord = currentWord
// 			}
// 			currentWord = "" // reset for next word
// 		}

// 	}

// 	// Check the last word (if no trailing space)
// 	if len(currentWord) > maxLen {
// 		maxLen = len(currentWord)
// 		longestWord = currentWord
// 	}

// 	fmt.Println("Longest word:", longestWord)
// 	fmt.Println("Length:", maxLen)

// }

func main() {
	str := "my name is Aditya"

	currentWord := ""
	longestWord := ""
	MaxLength := 0

	for _, ch := range str {

		if ch != ' ' {
			currentWord += string(ch)
		} else {
			if len(currentWord) > MaxLength {
				MaxLength = len(currentWord)
				longestWord = currentWord
			}
			currentWord = ""
		}

	}
	if len(currentWord) > MaxLength {
		MaxLength = len(currentWord)
		longestWord = currentWord
	}
	fmt.Println("Max Length is: ", MaxLength)
	fmt.Println("LongestWord is: ", longestWord)
}
