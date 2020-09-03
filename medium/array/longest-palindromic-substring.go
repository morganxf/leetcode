package array

// 最长回文子串
func longestPalindrome(s string) string {
	return LongestPalindrome_1(s)
}

func LongestPalindrome_1(s string) string {
	var length, index int
	for i := 0; i < len(s); i++ {
		for left, right := i, i; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] == s[right] {
				if left == 0 || right == len(s)-1 {
					if length < right-left+1 {
						index = left
						length = right - left + 1
					}
				}
				continue
			}
			if length < right-left-1 {
				index = left + 1
				length = right - left - 1
			}
			break
		}
	}
	for i := 0; i < len(s); i++ {
		for left, right := i, i+1; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] == s[right] {
				if left == 0 || right == len(s)-1 {
					if length < right-left+1 {
						index = left
						length = right - left + 1
					}
				}
				continue
			}
			if length < right-left-1 {
				index = left + 1
				length = right - left - 1
			}
			break
		}
	}
	return s[index : index+length]
}
