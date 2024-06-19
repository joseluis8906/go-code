package algorithms

func LongestPalindromicSubstr(str string) string {
	length := len(str)
	maxLen := 1
	start := 0

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			flag := true

			for k := 0; k < (j-i+1)/2; k++ {
				if str[i+k] != str[j-k] {
					flag = false
					break
				}
			}

			if flag && (j-i+1) > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	end := start + maxLen
	return str[start:end]
}
