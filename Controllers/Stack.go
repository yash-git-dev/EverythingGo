package Controllers

type Stack struct {
	items []int32
}

func (s *Stack) Push(item int32) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int32 {
	l := len(s.items)
	if l == 0 {
		return 0
	}
	removedItem := s.items[l-1]
	s.items = s.items[:l-1]
	return removedItem
}

func IsBalanced(input string) bool {
	stack := &Stack{}
	var openBrackets = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, val := range input {
		if _, ok := openBrackets[val]; ok {
			stack.Push(val)
		} else {
			if len(stack.items) != 0 && openBrackets[stack.Pop()] != val {
				return false
			}
		}
	}

	return len(stack.items) == 0
}
