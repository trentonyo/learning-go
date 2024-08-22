package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func testRemoveElement() {
	elementList := []int{0, 1, 2, 2, 3, 0, 4, 2}
	toRemove := 2
	expectedList := []int{0, 1, 4, 0, 3}

	k := removeElement(elementList, toRemove)

	rightSize := k == (len(elementList) - len(expectedList))
	rightItems := true
	sort.Ints(elementList)

	for i := 0; i < len(expectedList); i++ {
		if elementList[i] != expectedList[i] {
			rightItems = false
		}
	}
	fmt.Println("RightSize:", rightSize, "RightItems:", rightItems)
}

func removeDuplicates(nums []int) int {
	sticky := 0

	for _, elem := range nums[1:] {
		if nums[sticky] != elem {
			sticky++
			nums[sticky] = elem
		}
	}

	return sticky + 1
}

func testRemoveDuplicates() {
	input := []int{1, 1, 2}
	expectedUnique := 2

	k := removeDuplicates(input)
	fmt.Println("Test:", expectedUnique == k)
}

func removeDuplicates2(nums []int) int {
	k := 0
	offset := 0

	lastIndex, lastElement := 0, nums[0]
	for index, element := range nums[1:] {
		index++ // Indexes always start at 0

		// While within range of the last accepted value
		if index-lastIndex < 2 {
			// Move the window down if we find a new element
			if element != lastElement {
				lastElement = element
				lastIndex = index
			}

			// In either case, we are going to increment k. It is either valid because we are within range of the last accepted value or valid because we've discovered another unique value
			k = index - offset

			// If we have an offset, we need to shift values down
			if offset > 0 {
				nums[index-offset] = nums[index]
			}

			// If we are out of range and discover a new value, then we shift the window up
		} else if element != lastElement {
			// If we have an offset, we need to shift values down
			if offset > 0 {
				nums[index-offset] = nums[index]
			}

			lastElement = element
			lastIndex = index
			k = index - offset

			// We now have too many of one value
		} else {
			offset++
		}
	}

	return k + 1
}

func testRemoveDuplicates2() {
	//input := []int{0, 1, 2, 3}
	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expectedTotal := 7

	k := removeDuplicates2(input)
	fmt.Println("Test:", expectedTotal == k)
}

func majorityElement(nums []int) int {
	leader := nums[0]
	votes := 0

	for _, elem := range nums {
		if leader == elem {
			votes++
		} else if votes == 0 {
			leader = elem
		} else {
			votes--
		}
	}

	return leader
}

func rotate(nums []int, k int) {
	for r := 0; r < k; r++ {
		tmp := nums[len(nums)-1]

		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
	}
}

func testRotate() {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3

	rotate(nums, k)
}

type Tuple struct {
	First  int
	Second int
}

func stocks1(prices []int) int {
	profit := 0
	lowestPrice := prices[0]

	for _, price := range prices {
		if price < lowestPrice {
			lowestPrice = price
		} else if price-lowestPrice > profit {
			profit = price - lowestPrice
		}
	}

	return profit
}

func testStocks1() {
	six := stocks1([]int{1, 2, 3, 4, 5, 6, 7})
	zero := stocks1([]int{7, 6, 5, 4, 3, 2, 1})
	one := stocks1([]int{7, 6, 5, 4, 3, 2, 3})
	eight := stocks1([]int{3, 7, 2, 1, 5, 9, 6})
	println(six, zero, one, eight)

	print(stocks1([]int{886, 729, 539, 474, 5, 653, 588, 198, 313, 111, 38, 808, 848, 364, 819, 747, 520, 568, 583, 253, 605, 442, 779, 903, 217, 284, 927, 33, 859, 75, 418, 612, 174, 316, 167, 40, 945, 740, 174, 279, 985, 133, 38, 919, 528, 844, 101, 291, 673, 561, 244, 827, 602}))
}

func main() {
	//testRemoveElement()
	//testRemoveDuplicates()
	//testRemoveDuplicates2()
	//testRotate()
	testStocks1()
}
