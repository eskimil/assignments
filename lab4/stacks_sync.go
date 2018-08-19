// +build !solution

// Leave an empty line above this comment.
package lab4

type SafeStack struct {
	top  *Element
	size int
}

func (ss *SafeStack) Len() int {
	return ss.size
}

func (ss *SafeStack) Push(value interface{}) {
	ss.top = &Element{value, ss.top}
	ss.size++
}

func (ss *SafeStack) Pop() (value interface{}) {
	if ss.size > 0 {
		value, ss.top = ss.top.value, ss.top.next
		ss.size--
		return
	}

	return nil
}
