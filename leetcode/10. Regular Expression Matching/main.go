package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("mississippi", "mis*is*.p*."))
}

func isMatch(s string, p string) bool {
	atLeastOneCharPattern := "[a-z]+"

	pattern := ``

	for _, char := range p {
		switch char {
		case '*':
			lastChar := pattern[len(pattern)-1:]
			if lastChar == "+" {
				pattern += atLeastOneCharPattern
				break
			}
			pattern += lastChar
		case '.':
			pattern += atLeastOneCharPattern
		default:
			pattern += string(char)
		}
	}

	// fmt.Println("text", s, "pattern", pattern)

	re := regexp.MustCompile(fmt.Sprintf("^%s$", pattern))

	return re.Match([]byte(s))
}
