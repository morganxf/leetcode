package array

import "math"

func increasingTriplet(nums []int) bool {
	return IncreasingTriplet_1(nums)
}

func IncreasingTriplet_1(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	firstNum := math.MaxInt32
	secondNum := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= firstNum {
			firstNum = nums[i]
		} else if nums[i] <= secondNum {
			secondNum = nums[i]
		} else if nums[i] > secondNum {
			return true
		}
	}
	return false
}
