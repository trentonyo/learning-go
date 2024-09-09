package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// If it is okay to destroy the list (and it is not explicitly forbidden), then the quickest
	//  way is to create a "seen" node and after each step on the list set its Next to that node.
	//  That way, if any node points to that one then we know we've seen if before

	// Edge case, empty lists don't have cycles
	if head == nil {
		return false
	}

	seenNode := new(ListNode)

	current := head
	var last *ListNode

	for current.Next != nil {
		// Iff we have seen this node, there is a cycle
		if current.Next == seenNode {
			return true
		}

		// Move to the next link
		last = current
		current = current.Next

		// Set the last link's Next to seen
		last.Next = seenNode
	}

	// Iff we have made it to the end of the list without seeing a node twice then there is no cycle
	return false
}

// func testHasCycle() {

// 	fmt.Printf("%v := %v\n", hasCycle("]"), false)
// }

func main() {
	hasCycle(nil)
}
