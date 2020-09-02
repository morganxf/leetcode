package array

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	//return LengthOfLongestSubstring_1(s)
	return LengthOfLongestSubstring_2(s)
}

func LengthOfLongestSubstring_1(s string) int {
	var maxLen int
	for i := 0; i < len(s); i++ {
		tmp := make(map[byte]int)
		for j := i; j < len(s); j++ {
			if index, found := tmp[s[j]]; !found {
				tmp[s[j]] = j
				if j == len(s)-1 {
					i = len(s)
				}
			} else {
				i = index
				break
			}
		}
		if maxLen < len(tmp) {
			maxLen = len(tmp)
		}
	}
	return maxLen
}

func LengthOfLongestSubstring_2(s string) int {
	var maxLen int
	var start int = 0
	tmp := make(map[byte]int, len(s))
	for i := 0; i < len(s); {
		for j := i; j < len(s); j++ {
			index, found := tmp[s[j]]
			tmp[s[j]] = j
			if j == len(s)-1 {
				var delta int = 1
				if found && index >= start {
					delta = 0
				}
				if maxLen < j-start+delta {
					maxLen = j - start + delta
				}
				i = j + 1
				break
			}
			if found && index >= start {
				if maxLen < j-start {
					maxLen = j - start
				}
				start = index + 1
				i = j + 1
				break
			}
		}
	}
	return maxLen
}
