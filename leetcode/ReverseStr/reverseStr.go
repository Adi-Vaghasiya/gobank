package main

import (
	"fmt"
	"sort"
)

//	func ReverseStr(s string) string {
//		var left, right int
//		convertedByteString := []byte(s)
//		Lenstring := len(convertedByteString)
//		mid := Lenstring / 2
//		if Lenstring%2 == 0 {
//			left = mid - 1
//			right = mid
//		} else {
//			left = mid - 1
//			right = mid + 1
//		}
//		for left >= 0 && right < Lenstring {
//			convertedByteString[left], convertedByteString[right] = convertedByteString[right], convertedByteString[left]
//			left--
//			right++
//		}
//		return string(convertedByteString)
//	}
func sorting(nums []int) []int {
	// number := make([]int, len(nums))
	// copy(number, nums)
	fmt.Println("length of Nums is: ", len(nums))
	// SortedArray := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= nums[i] {
				// tmp := number[j]
				// number[i] = tmp
				nums[j], nums[i] = nums[i], nums[j]
				// SortedArray = append(SortedArray, number[i])
				sort.SliceStable(nums, func(i, j int) bool {
					return true
				})

			}
		}
	}
	// return SortedArray
	return nums
}

func main() {
	nums := []int{1, 8, 2, 5, 6, 7, 3}
	// Str := "121"
	//dlrowolleh
	// recivedStr := ReverseStr(Str)
	SortedArray := sorting(nums)
	fmt.Println(SortedArray)
}
