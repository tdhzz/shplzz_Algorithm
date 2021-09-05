package yan

import "strings"

func lengthOfLastWord(s string) int {
	str := strings.TrimSpace(s)
	if str == "" {
		return 0
	}

	length := 0
	for i := len(str); i > 0; i-- {
		currStr := str[i-1:i]
		if currStr == " " {
			break
		}
		length++
	}
	return length
}