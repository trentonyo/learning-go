package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	// Edge case, a ransomNote longer than the magazine is impossible
	if len(ransomNote) > len(magazine) {
		return false
	}

	noteIdx := 0
	earlyExit := true
	// Consume each character that is used from magazine and ransomeNote
	for (len(ransomNote) > 0 || len(magazine) > 0) && noteIdx < len(ransomNote) {
		// Iterate through the note, because if any character of the note is not found in the magazine then the construction is impossible (with an index, to maintain control flow)
		noteChar := ransomNote[noteIdx]

		// Iterate through the magazine, exiting and removing it when this letter of the note is found
		// If the char is never found, then there is no solution and earlyExit falls through
		for magazineIdx, magazineChar := range magazine {
			if noteChar == byte(magazineChar) {
				// Remove char from needed chars to construct ransomNote, adjusting the index
				ransomNote = ransomNote[:noteIdx] + ransomNote[noteIdx+1:]
				noteIdx--

				// Cut char from the magazine
				magazine = magazine[:magazineIdx] + magazine[magazineIdx+1:]

				earlyExit = false
				break
			}
		}

		// When a char of the note is not found in the magazine, the construct fails
		if earlyExit {
			return false
		}

		noteIdx++
	}

	return len(ransomNote) == 0
}

func testCanConstruct() {
	fmt.Printf("%t := %t\n", canConstruct("a", "b"), false)
	fmt.Printf("%t := %t\n", canConstruct("aa", "ab"), false)
	fmt.Printf("%t := %t\n", canConstruct("aa", "aab"), true)
	fmt.Printf("%t := %t\n", canConstruct("bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"), true)
}

func isIsomorphic(s string, t string) bool {
	// Use a hashmap to drive the lookup
	var sToT = make(map[string]string)
	var tToS = make(map[string]string)

	// Iterate through s (len(s) == len(t) so the outer loop is arbitrary)
	for i := range s {
		sChar := s[i : i+1]
		tChar := t[i : i+1]

		if sToT[sChar] == "" { // If s does not yet have a mapping to t, set one
			sToT[sChar] = tChar
		} else if sToT[sChar] != tChar { // If s to t is mapped and this violates, then fail
			return false
		}

		if tToS[tChar] == "" { // If t does not yet have a mapping to s, set one
			tToS[tChar] = sChar
		} else if tToS[tChar] != sChar { // If t to s is mapped and this violates, then fail
			return false
		}
	}

	// If we made it all the way to here, the strings are isomorphic
	return true
}

func testIsIsomorphic() {
	fmt.Printf("%t := %t\n", isIsomorphic("a", "b"), true)
	fmt.Printf("%t := %t\n", isIsomorphic("foo", "bar"), false)
	fmt.Printf("%t := %t\n", isIsomorphic("badc", "baba"), false)
	fmt.Printf("%t := %t\n", isIsomorphic("paper", "title"), true)
}

func wordPattern(pattern string, s string) bool {
	// Use a hashmap to drive the lookup
	var patternToWord = make(map[string]string)
	var wordToPattern = make(map[string]string)
	patternIdx := 0

	wordStart, wordEnd := 0, 0

	// Iterate through the string, collecting words and comparing to pattern
	for i := range s {
		sChar := s[i : i+1]

		// If we've reached a space or the end of the string...
		if sChar == " " || i == len(s)-1 {
			// (special case: end of string needs to capture last char)
			if i == len(s)-1 {
				wordEnd++
			}

			// ...start a new word, and...
			word := s[wordStart:wordEnd]
			wordStart, wordEnd = i+1, i+1

			// ...process the word found
			pKey := pattern[patternIdx : patternIdx+1]

			if patternToWord[pKey] == "" && wordToPattern[word] == "" { // If the pattern key isn't in the map yet, set it to this word
				patternToWord[pKey] = word
				wordToPattern[word] = pKey
			} else if patternToWord[pKey] != word || wordToPattern[word] != pKey { // If the pattern is broken, fail
				return false
			}

			patternIdx++

			// Check if pattern is not long enough for string
			if patternIdx >= len(pattern) && wordEnd < len(s) {
				return false
			}
		} else {
			wordEnd++
		}
	}

	// If we made it all the way to here and the pattern is exhausted
	return patternIdx == len(pattern)
}

func testWordPattern() {
	fmt.Printf("%t := %t\n", wordPattern("abba", "dog cat cat dog"), true)
	fmt.Printf("%t := %t\n", wordPattern("abba", "dog dog dog dog"), false)
	fmt.Printf("%t := %t\n", wordPattern("abc", "dog cat dog"), false)
	fmt.Printf("%t := %t\n", wordPattern("aaa", "aa aa aa aa"), false)
}

func isAnagram(s string, t string) bool {
	// First pass, count the incidence of all chars of s
	var charIncidence = make(map[string]int)
	individualChars := 0
	completedChars := 0

	for i := range s {
		if charIncidence[s[i:i+1]] == 0 {
			individualChars++
		}
		charIncidence[s[i:i+1]]++
	}

	// Second pass, decrement the incidence of all chars of s (in the same map)...
	for i := range t {
		charIncidence[t[i:i+1]]--

		if charIncidence[t[i:i+1]] < 0 { // ...fail if incidence < 0 after decrement
			return false
		} else if charIncidence[t[i:i+1]] == 0 { // Track chars that have been completely accounted for
			completedChars++
		}
	}

	// Make sure that each chars incidence is 0
	return completedChars == individualChars
}

func testIsAnagram() {
	fmt.Printf("%t := %t\n", isAnagram("anagram", "nagaram"), true)
	fmt.Printf("%t := %t\n", isAnagram("anagram", "gross face"), false)
}

func twoSum(nums []int, target int) []int {
	var indexByVal = make(map[int][]int)

	// First pass, store all indices where this value occurs
	for i, val := range nums {
		indexByVal[val] = append(indexByVal[val], i)
	}

	// Second pass, look for the needed values
	solution := []int{}

	for iL, val := range nums {
		// ONLY consider values less than target-1, this disregards numbers that can't possibly result in a two-value solution
		if target-val > 0 || target <= 1 { // EDGE CASE: if the target necessarily requires a value to be zero
			needed := target - val

			// considering the list of indices that contain the needed value...
			for _, iR := range indexByVal[needed] {
				// ..find the index which is different from the current one
				if iL != iR {
					solution = []int{iL, iR}
				}
			}
		}
	}

	return solution
}

func testTwoSum() {
	fmt.Printf("%v := %v\n", twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3})
	fmt.Printf("%v := %v\n", twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	fmt.Printf("%v := %v\n", twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	fmt.Printf("%v := %v\n", twoSum([]int{3, 3}, 6), []int{1, 0})
}

func getDigits(n int) []int {
	output := []int{}
	principle := n

	for n > 0 {
		principle = (n / 10) * 10
		digit := n - principle
		n = int(n / 10)

		// The order of the digits is reversed, but due to the transitive property of addition that is okay for THIS problem
		output = append(output, digit)
	}
	return output
}

func calculateHappy(n int) int {
	digits := getDigits(n)
	newSum := 0

	for _, val := range digits {
		newSum += val * val
	}

	return newSum
}

func isHappy(n int) bool {
	var seenSums = make(map[int]bool)

	it := n
	for {
		// Caluculate the digits
		it = calculateHappy(it)

		// If a cycle is detected, fail
		if seenSums[it] {
			return false
		}

		// If the number is 1, succeed
		if it == 1 {
			return true
		}

		// Mark this sum as having been seen
		seenSums[it] = true
	}
}

func testIsHappy() {
	fmt.Printf("%t := %t\n", isHappy(19), true)
	fmt.Printf("%t := %t\n", isHappy(2), false)
}

func main() {
	testIsHappy()
}
