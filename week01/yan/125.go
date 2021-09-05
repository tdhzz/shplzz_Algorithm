package yan

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	rep := regexp.MustCompile(`[^a-zA-Z]+`)
	str := rep.ReplaceAllString(s, "")
	normalStr := strings.ToLower(str)

	reverseStr := strReverse(normalStr)
	if normalStr != reverseStr {
		return false
	}
	return true
}

func strReverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	str := ""
	for i:=len(s); i > 0; i-- {
		str += s[i-1:i]
	}
	return str
}
