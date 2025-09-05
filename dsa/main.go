package main

import "fmt"

func main() {
	var array = []int{1, 3, 5, 6}
	var result = searchInsert(array, 5)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	//your code goes here
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var low = 0
	var high = len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if nums[high] > target {
		return high - 1
	}
	return high + 1
}

func getFloorAndCeil(nums []int, x int) []int {
	if x < nums[0] {
		return []int{-1, nums[0]}
	}
	if x > nums[len(nums)-1] {
		return []int{nums[len(nums)-1], -1}
	}
	var low = 0
	var high = len(nums) - 1
	var result = make([]int, 2)

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == x {
			return []int{nums[mid], nums[mid]}
		} else if nums[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if nums[high] > x {
		result[0] = nums[high-1]
		result[1] = nums[high]
	} else {
		result[0] = nums[high]
		result[1] = nums[high+1]
	}
	return result
}
