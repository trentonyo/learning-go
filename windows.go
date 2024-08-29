package main

import (
	"fmt"
	"math"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	minSubLen := math.MaxInt

	windowSum := 0
	l, r := 0, 0

	for r <= len(nums) {
		if windowSum >= target {
			minSubLen = min(r-l, minSubLen)

			windowSum -= nums[l]
			l++

			continue
		}

		if r < len(nums) {
			r++
		} else {
			break
		}
		windowSum += nums[r-1]
	}

	if minSubLen == math.MaxInt {
		minSubLen = 0
	}

	return minSubLen
}

func testMinSubArrayLen() {
	fmt.Printf("%d := %d\n", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
	fmt.Printf("%d := %d\n", minSubArrayLen(4, []int{1, 4, 4}), 1)
	fmt.Printf("%d := %d\n", minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1}), 0)
}

func hasRepeats(s string) bool {
	// This is terrible
	done := ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(done, s[i:i+1]) {
			return true
		} else {
			done += s[i : i+1]
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxSubLen := 0

	windowSlice := ""
	l, r := 0, 1

	windowSlice = s[l:r]
	for r <= len(s) {
		// This approach would probably be much better if it did not check the whole slice every time.
		if hasRepeats(windowSlice) && l < r {
			l++
			windowSlice = s[l:r]

			continue
		} else {
			maxSubLen = max(len(windowSlice), maxSubLen)
		}

		if r < len(s) {
			r++
		} else {
			break
		}
		windowSlice = s[l:r]
	}

	if maxSubLen == math.MaxInt {
		maxSubLen = 0
	}

	return maxSubLen
}

func testLengthOfLongestSubstring() {
	fmt.Printf("%d := %d\n", lengthOfLongestSubstring(" "), 1)
	fmt.Printf("%d := %d\n", lengthOfLongestSubstring("abcabcbb"), 3)
	fmt.Printf("%d := %d\n", lengthOfLongestSubstring("bbbbbbb"), 1)
	fmt.Printf("%d := %d\n", lengthOfLongestSubstring("pwwkekekekewkaewkeww"), 4)
}

func main() {
	testLengthOfLongestSubstring()
}
