package array

import "sort"

func threeSum(nums []int) [][]int {
	//return Sum3Nums_1(nums)
	return Sum3Nums_2(nums)
}

// first < second < third
func Sum3Nums_2(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := len(nums) - 1
		target := -nums[first]
		for second < third {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if nums[second]+nums[third] < target {
				second++
			} else if nums[second]+nums[third] > target {
				third--
			} else {
				result = append(result, []int{nums[first], nums[second], nums[third]})
				second++
				third--
			}
		}
	}
	return result
}

func Sum3Nums_1(nums []int) [][]int {
	result := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		var numsTmp []int
		for j, v := range nums {
			if i == j {
				continue
			}
			numsTmp = append(numsTmp, v)
		}
		var sum2 int = 0 - nums[i]
		result2Num := Sum2Nums(numsTmp, sum2)
		for _, item := range result2Num {
			if nums[i] < item[0] {
				result[[3]int{nums[i], item[0], item[1]}] = true
			} else if nums[i] < item[1] {
				result[[3]int{item[0], nums[i], item[1]}] = true
			} else {
				result[[3]int{item[0], item[1], nums[i]}] = true
			}
		}
	}
	tmp := make([][]int, 0, len(result))
	for key := range result {
		tmp = append(tmp, []int{key[0], key[1], key[2]})
	}
	return tmp
}

type twoNum struct {
	numsSet  map[[2]int]bool
	numsList [][2]int
}

func NewTwoNum() *twoNum {
	tn := &twoNum{
		numsSet: make(map[[2]int]bool),
	}
	return tn
}

func (tn *twoNum) AddToSet(v [2]int) {
	tn.numsSet[v] = true
}

func (tn *twoNum) AddToSetWithSortedNums(v [2]int) {
	if v[0] > v[1] {
		v[0], v[1] = v[1], v[0]
	}
	tn.AddToSet(v)
}

func (tn *twoNum) ToList() [][2]int {
	if len(tn.numsList) > 0 {
		return tn.numsList
	}
	var result [][2]int
	for key := range tn.numsSet {
		result = append(result, key)
	}
	return result
}

func Sum2Nums(nums []int, sum int) [][2]int {
	result := NewTwoNum()
	difference := make(map[int]bool)
	for _, v := range nums {
		if ok := difference[sum-v]; ok {
			result.AddToSetWithSortedNums([2]int{sum - v, v})
		} else {
			difference[v] = true
		}
	}
	return result.ToList()
}
