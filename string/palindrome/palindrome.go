package main

import "fmt"

func PalindromeStr(s string) int {
	// l := 0
	// r := 0
	count := 0
	for center := 0; center < len(s); center++ {

		l, r := center, center
		for l >= 0 && r < len(s) && s[l] == s[r] {
			fmt.Printf("Found odd-length palindrome: %s \n", s[l:r+1])
			count++
			l--
			r++
		}
		l, r = center, center+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			fmt.Printf("Found even-length palindrome:: %s \n", s[l:r+1])
			count++
			l--
			r++
		}

	}
	return count
}

func main() {
	count := PalindromeStr("aabbaabbaa")
	fmt.Printf("The Count is: %v\n", count)
}

// r = i + 1
// check := s[i] == s[r]

// if check == true {

// 	i++
// 	// if i == 1 {
// 	// 	check1 := s[l] == s[r]

// 	count = r - l + 1
// 	if i%2 == 1 && s[l] == s[r] {
// 		l = i - 1

// 		r++
// 		fmt.Println("Increased R and Decresed L by 1")
// 		fmt.Printf("found length of %v Palindrome: %c%c%c \n", count, s[l], s[i], s[r])
// 	}
// 	// }
