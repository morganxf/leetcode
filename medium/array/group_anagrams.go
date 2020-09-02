package array

// 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	return GroupAnagrams(strs)
}

type StrIndex struct {
	index [26]int
}

// hash function
func NewStrIndex(str string) StrIndex {
	var strIndex StrIndex
	for i := 0; i < len(str); i++ {
		strIndex.index[str[i]-'a']++
	}
	return strIndex
}

func GroupAnagrams(strs []string) [][]string {
	resultMap := make(map[StrIndex][]string)
	for _, str := range strs {
		strIndex := NewStrIndex(str)
		resultMap[strIndex] = append(resultMap[strIndex], str)
	}
	var result [][]string
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
