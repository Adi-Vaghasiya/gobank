package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			indexes := []int{
	// 				i, j,
	// 			}
	// 			return indexes
	// 		}

	// 	}
	// }
	Hashmap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, found := Hashmap[diff]; found {
			return []int{j, i}
		}
		Hashmap[num] = i
		fmt.Println(Hashmap)
	}
	return nil
}

func main() {
	target := 9
	intArry := []int{
		7, 8, 2, 15,
	}
	indexes := twoSum(intArry, target)
	fmt.Println("Indexes that are add up to Target are: ", indexes)
}
