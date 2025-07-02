// package main

// import "fmt"

// // func longestSubsequenceRepeatedK(s string, k int) []string {
// // count := 0
// // Slice := []string{}
// // for i := 0; i < len(s); i++ {
// // for j := i + 1; j < len(s); j++ {
// // if s[j] == s[i] {
// // count++
// // Slice = (append(Slice, s[i:j+1]))
// // break
// // }
// // }
// // }
// // // fmt.Println(hashmap)
// // return Slice
// // }

// // func ComparingSubsequence(dividedstrings []string) ([]string, int) {
// // repeats := make(map[string]int)
// // seen := make(map[string]bool)

// // for i := 0; i < len(dividedstrings); i++ {
// // for j := i + 1; j < len(dividedstrings); j++ {
// // str1 := dividedstrings[i]
// // str2 := dividedstrings[j]

// // if len(str1) != len(str2) {
// // continue
// // }
// // if str1 == "" || str2 == "" {
// // continue
// // }

// // match := true
// // for x := 0; x < len(str1); x++ {
// // if str1[x] != str2[x] {
// // match = false
// // break
// // }
// // }

// // if match {
// // repeats[str1]++
// // }
// // }
// // }

// // // Collect repeated sequences (occurred more than once)
// // result := []string{}
// // maxCount := 0

// // for seq, count := range repeats {
// // if count >= 1 {
// // if !seen[seq] {
// // result = append(result, seq)
// // seen[seq] = true
// // }
// // if count > maxCount {
// // maxCount = count
// // }
// // }
// // }

// // return result, maxCount + 1
// // }

// // func ComparingSubsequence(dividedstrings []string) ([]string, int) {
// // for i := 0; i < len(dividedstrings); i++ {
// // for j := i + 1; j < len(dividedstrings); j++ {
// // x, y, count := 0, 0, 0
// // for x < len(dividedstrings[i]) && y < len(dividedstrings[j]) {
// // if dividedstrings[i][x] == dividedstrings[j][y] {
// // count++
// // x++
// // y++
// // }

// // }

// // // fmt.Println(dividedstrings[i], " ", dividedstrings[j])
// // }
// // }
// // return nil, 2
// // }

// // func longestSubsequenceRepeatedK(s string, k int) []string {
// // count := 0
// // Slice := make([]string, len(s))
// // for i := 0; i < len(s); i++ {
// // for j := i + 1; j < len(s); j++ {
// // if s[j] == s[i] {
// // count++
// // Slice[count-1] = s[i : j+1]
// // break
// // }
// // }
// // }
// // // fmt.Println(hashmap)
// // return Slice
// // }
// func longestSubsequenceRepeatedK(s string, k int) ([]string, int) {
// 	RepeatedStr := []string{}
// 	result := ""
// 	count := 1
// 	for i := 0; i < len(s); i++ {
// 		found := false
// 		fmt.Println("This is I: ", i)
// 		for j := i + 1; j < len(s); j++ {
// 			fmt.Println("This is J: ", j)
// 			if s[i] == s[j] {
// 				count++
// 				result += string(s[i])
// 				found = true
// 				break
// 				//letsleetcode
// 			}

// 		}
// 		if !found && result != "" {
// 			RepeatedStr = append(RepeatedStr, result)
// 			result = ""
// 			// count = 1
// 			// Optional: do something if there's no match
// 			// Like: results = append(results, "")
// 		}

// 	}

// 	return RepeatedStr, len(RepeatedStr)
// }

// func main() {
// 	stringis := "letsleetcode"
// 	initial_k := 0
// 	RecivedStr, k := longestSubsequenceRepeatedK(stringis, initial_k)
// 	fmt.Printf("The repeated Seq is: %v and it occurs %v times", RecivedStr, k)
// 	// Repeated, Max := ComparingSubsequence(RecivedStr)
// 	// fmt.Println("Maximum repetition count:", Max)
// 	// fmt.Println("Repeated sequences:", Repeated)
// 	//fmt.Println(RecivedStr)

// }

// //letsleetcode
package main

import (
	"fmt"
)

func isSubsequence(sub, full string, k int) bool {
	count := 0
	idx := 0
	//letsleetcode
	for i := 0; i < len(full); i++ {
		if idx < len(sub) && full[i] == sub[idx] {
			idx++
			if idx == len(sub) {
				count++
				if count >= k {
					return true
				}
				idx = 0 // restart to look for another occurrence
			}
		}
	}
	return false
}

func generateSubsequences(s string, path string, index int, k int, results map[string]bool) {
	if len(path) > 0 && isSubsequence(path, s, k) {
		results[path] = true
	}
	if index == len(s) {
		return
	}
	// include s[index]
	generateSubsequences(s, path+string(s[index]), index+1, k, results)
	// exclude s[index]
	generateSubsequences(s, path, index+1, k, results)

}

// func CheckFunc(s []string) string {
// 	return s[1]
// }

func main() {
	longest := ""
	s := "letsleetcode"
	k := 2
	results := make(map[string]bool)
	generateSubsequences(s, "", 0, k, results)

	// longest := ""
	for str := range results {
		if len(str) > len(longest) || (len(str) == len(longest) && str > longest) {
			longest = str
			// longest = append(longest, str)
		}
	}
	// Str := CheckFunc(longest)
	// fmt.Println(Str)
	fmt.Println("All repeated subsequences (k ≥ 2):")
	for str := range results {
		fmt.Println(" -", str)
	}
	fmt.Println("Longest:", longest)

}
