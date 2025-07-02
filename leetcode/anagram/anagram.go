// package main

// import "fmt"

// func anagramCheking(s []string) [][]string {
// 	// data := make([][]string, len(s))
// 	fmt.Println(s)
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {

// 			if len(s[i]) != len(s[j]) {
// 				fmt.Println(j)
// 				continue

// 			}
// 			x, y := 0, 0
// 			for x < len(s[i]) && y < len(s[j]) && s[i][x] == s[j][y] {
// 				fmt.Printf("Match: %c == %c\n", s[i][x], s[j][y])
// 				x++
// 				y++
// 			}

// 			// if len(s[i]) == len(s[j]) {

// 			// }

// 		}
// 	}
// 	return nil
// }

//	func main() {
//		s := []string{
//			"eat", "teaaa", "tan", "ate", "nat", "bat",
//		}
//		anagramCheking(s)
//	}
package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	r := []rune(s)
	sort.SliceStable(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sorted := sortString(word)
		// fmt.Println(sorted)
		anagramMap[sorted] = append(anagramMap[sorted], word)

	}
	// fmt.Println(anagramMap)
	// Collect results
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
		// fmt.Println(result)
	}
	return result
}

// func groupAnagramsss(strs []string) [][]string {
// 	anagramMap := make(map[string][]string)

// 	for _, word := range strs {
// 		// Frequency key as string (e.g., "1#0#0#0#..." for "a")
// 		count := make([]int, 26)
// 		for _, ch := range word {
// 			count[ch-'a']++
// 		}

// 		// Turn count array into a string key
// 		key := make([]byte, 0, 52)
// 		for _, c := range count {
// 			key = append(key, byte(c)+'0', '#') // '#' as separator
// 		}

// 		anagramMap[string(key)] = append(anagramMap[string(key)], word)
// 	}

// 	// Collect results
// 	result := make([][]string, 0, len(anagramMap))
// 	for _, group := range anagramMap {
// 		result = append(result, group)
// 	}
// 	return result
// }

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(s)
	// fmt.Println(result[0])
	fmt.Println(result)
}
