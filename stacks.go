package main

import (
	"fmt"
	"strings"
)

var BRACKETS = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

//type StackMaker interface {
//	Push(v string)
//	Pop() string
//}

type StringStack struct {
	internal []string
}

func (x *StringStack) Push(v string) {
	x.internal = append(x.internal, v)
}

func (x *StringStack) Pop() string {
	output := x.internal[len(x.internal)-1]
	x.internal = x.internal[:len(x.internal)-1]
	return output
}

func (x *StringStack) isEmpty() bool {
	return len(x.internal) == 0
}

func (x *StringStack) getInternal() []string {
	return x.internal
}

func simplifyPath(path string) string {
	stack := new(StringStack)
	path += "/"

	word := ""

	// If path doesn't start with a slash, add it
	if path[:1] != "/" {
		path = "/" + path
	}

	// Collect words and add them to the stack
	l, r := 0, 0

	for r < len(path) {
		r++
		if (r+1 < len(path) && path[r:r+1] == "/") || r+1 == len(path) {
			if r+1 == len(path) {
				r++
			}

			// If the next char is a slash as well, just continue, a word hasn't started
			if r+2 < len(path) && path[r+1:r+2] == "/" {
				continue
			}

			word = path[l:r]

			// Handle word

			// Truncate trailing slashes
			t := len(word) - 1
			for word[t:t+1] == "/" && t > 0 {
				word = word[:t]
				t--
			}

			// "/" is invisible
			// "/." is invisible
			// "//" is invisible
			if word == "/" || word == "/." || word == "//" {

			} else if word == "/.." { // "/.."" pops the stack
				if !stack.isEmpty() {
					stack.Pop()
				}
			} else { // everything else is pushed to the stack
				stack.Push(word)
			}

			l = r
			r = r + 1
		}
	}

	// Edge case, root
	if stack.isEmpty() {
		return "/"
	}

	return strings.Join(stack.getInternal(), "")
}

func testSimplifyPath() {
	fmt.Printf("%s := %s\n", simplifyPath("/.././GVzvE/./xBjU///../..///././//////T/../../.././zu/q/e"), "/zu/q/e")
	fmt.Printf("%s := %s\n", simplifyPath("///TJbrd/owxdG//"), "/TJbrd/owxdG")
	fmt.Printf("%s := %s\n", simplifyPath("///"), "/")
	fmt.Printf("%s := %s\n", simplifyPath("/../"), "/")
	fmt.Printf("%s := %s\n", simplifyPath("/home/user/Documents/../Pictures"), "/home/user/Pictures")
	fmt.Printf("%s := %s\n", simplifyPath("/home/"), "/home")
	fmt.Printf("%s := %s\n", simplifyPath("/home//////////out//"), "/home/out")
}

func isValid(s string) bool {
	stack := new(StringStack)

	// For each char in the string...
	for i := range s {
		v := s[i : i+1]

		//...if it is in BRACKETS, push the val of its complementary bracket
		closingVal, opening := BRACKETS[v]

		if opening {
			stack.Push(closingVal)
		} else {

			// Edge case, stack is empty when a closing bracket is encountered
			if stack.isEmpty() {
				return false
			}

			//...if v is a closing bracket, make sure it is the right one by popping from the stack
			if v != stack.Pop() {
				return false
			}
		}
	}

	return stack.isEmpty()
}

func testIsValid() {
	fmt.Printf("%v := %v\n", isValid("]"), false)
	fmt.Printf("%v := %v\n", isValid("()"), true)
	fmt.Printf("%v := %v\n", isValid("()[]{}"), true)
	fmt.Printf("%v := %v\n", isValid("(]"), false)
	fmt.Printf("%v := %v\n", isValid("([])"), true)
}

// func main() {
// 	testSimplifyPath()
// }
