package main

import (
	"fmt"
)

func rn(representation string) int {
	switch representation {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	sum := 0
	chunk := 0

	for i := 0; i < len(s); i++ {
		curr := rn(s[i : i+1])

		next := 0
		if i < len(s)-1 {
			next = rn(s[i+1 : i+2])
		}

		if curr < next {
			chunk -= curr
		} else {
			sum += curr + chunk
			chunk = 0
		}
	}

	return sum
}

func testRomanToInt() {
	fmt.Printf("%d := %d\n", romanToInt("III"), 3)
	fmt.Printf("%d := %d\n", romanToInt("LVIII"), 58)
	fmt.Printf("%d := %d\n", romanToInt("MCMXCIV"), 1994)
}

func intToRoman(num int) string {
	result := ""

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			result += symbols[i]
			num -= vals[i]
		}
	}

	return result
}

func testIntToRoman() {
	fmt.Printf("%s := %s\n", intToRoman(3), "III")
	fmt.Printf("%s := %s\n", intToRoman(58), "LVIII")
	fmt.Printf("%s := %s\n", intToRoman(1994), "MCMXCIV")
}

func lengthOfLastWord(s string) int {
	running := 0
	inSpace := false

	for i := 0; i < len(s); i++ {
		c := s[i : i+1]

		if c == " " {
			inSpace = true
		} else if inSpace {
			inSpace = false
			running = 1
		} else {
			running++
		}
	}

	return running
}

func testLengthOfLastWord() {
	fmt.Printf("%d := %d\n", lengthOfLastWord("luffy is still joyboy"), 6)
	fmt.Printf("%d := %d\n", lengthOfLastWord("   fly me   to   the moon  "), 4)
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// Starting at the end of this string or the prefix, whichever is shorter...
		earlyExit := false
		for j := min(len(strs[i]), len(prefix)); j > 0 && !earlyExit; j-- {
			// ...check if the prefix matches yet
			if strs[i][:j] == prefix[:j] {
				prefix = strs[0][:j]
				earlyExit = true
			}
		}

		if !earlyExit {
			return ""
		}
	}

	return prefix
}

func testLongestCommonPrefix() {
	fmt.Printf("%s := %s\n", longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	fmt.Printf("%s := %s\n", longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
}

func reverseWords(s string) string {
	output := ""

	start := 0
	wordStarted := false

	for i := 0; i < len(s)+1; i++ {
		c := " "
		if i < len(s) {
			c = s[i : i+1]
		}

		if c != " " {
			if !wordStarted {
				wordStarted = true
				start = i
			}
		} else if wordStarted {
			wordStarted = false
			word := s[start:i]

			if len(output) == 0 {
				output = word
			} else {
				output = word + " " + output
			}
		}
	}

	return output
}

func testReverseWords() {
	fmt.Printf("%s := %s\n", reverseWords("the sky is blue"), "blue is sky the")
	fmt.Printf("%s := %s\n", reverseWords("  hello world  "), "world hello")
}

func strStr(haystack string, needle string) int {
	index := -1
	match := 0

	if len(needle) > len(haystack) { // edge case: needle too long for haystack
		return index
	} else if len(needle) == 0 && len(haystack) == 0 { // edge case: empty strings
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		a := haystack[i]
		b := needle[match]

		if a == b {
			if match == 0 {
				index = i
			}
			match++

			if match == len(needle) {
				break
			}
		} else {
			if match > 0 {
				i = index // Start the search again from the start of this substring that failed
			}
			index = -1
			match = 0
		}
	}

	if match == len(needle) { // Check if we have a complete substr by the end, otherwise fail
		return index
	} else {
		return -1
	}
}

func testStrStr() {
	// fmt.Printf("%d := %d\n", strStr("sadbutsad", "sad"), 0)
	// fmt.Printf("%d := %d\n", strStr("leetcode", "leeto"), -1)
	// fmt.Printf("%d := %d\n", strStr("makin banana pancakes", "an"), 7)
	fmt.Printf("%d := %d\n", strStr("mississippi", "issip"), 4)
}

func main() {
	//testRomanToInt()
	//testLengthOfLastWord()
	//testLongestCommonPrefix()
	// testReverseWords()
	testStrStr()
}
