package flatten

import "slices"

func Flatten(nested interface{}) []interface{} {
	stack := []interface{}{nested}
	result := []interface{}{}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		switch t := curr.(type) {
		case []interface{}:
			stack = append(stack, t...)
		default:
			if t != nil {
				result = append(result, t)
			}
		}
	}
	slices.Reverse(result)
	return result
}
