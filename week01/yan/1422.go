package yan

func maxScore(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxScore := 0
	count := len(s)
	for i := 1; i < count; i++ {
		left := s[0:i]
		right := s[i:count]
		leftNum := searchKeywordNum(left, "0")
		rightNum := searchKeywordNum(right, "1")
		if leftNum + rightNum > maxScore {
			maxScore = leftNum + rightNum
		}
	}
	return maxScore
}

func searchKeywordNum(s, keyword string) int {
	if len(s) == 0 {
		return 0
	}

	num := 0
	count := len(s)
	for i := 0; i < count; i++ {
		if s[i:i+1] == keyword {
			num++
		}
	}
	return num
}