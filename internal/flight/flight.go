package flight

import "fmt"

func Parser(entry [][]string) ([]string, error) {
	answer := []string{}

	stack1 := []string{}
	stack2 := []string{}

	// for each entry we are going to push the element to the stack,
	// and if we find it in the other stack, we are going to pop it in both stacks
	for _, v := range entry {
		stack1 = append(stack1, v[0])
		stack2 = append(stack2, v[1])
		for i1, v1 := range stack1 {
			if v1 == v[1] {
				stack1 = append(stack1[:i1], stack1[i1+1:]...)
				stack2 = stack2[:len(stack2)-1]
				break
			}
		}
		for i2, v2 := range stack2 {
			if v2 == v[0] {
				stack2 = append(stack2[:i2], stack2[i2+1:]...)
				stack1 = stack1[:len(stack1)-1]
				break
			}
		}
	}

	// we can only have one in each stack, otherwise the entry we received is invalid
	if (len(stack1) > 1) || (len(stack2) > 1) {
		return answer, fmt.Errorf("invalid entry")
	}

	// fill the answer with the lefting result from the popping pushing above
	answer = append(answer, stack1[0], stack2[0])

	return answer, nil
}

func ParserConcurrency(entry [][]string) ([]string, error) {
	answer := []string{}

	stack1 := []string{}
	stack2 := []string{}

	// for each entry we are going to push the element to the stack,
	// and if we find it in the other stack, we are going to pop it in both stacks
	for _, v := range entry {
		stack1 = append(stack1, v[0])
		stack2 = append(stack2, v[1])
		go func() {
			for i1, v1 := range stack1 {
				if v1 == v[1] {
					stack1 = append(stack1[:i1], stack1[i1+1:]...)
					stack2 = stack2[:len(stack2)-1]
					break
				}
			}
		}()
		go func() {
			for i2, v2 := range stack2 {
				if v2 == v[0] {
					stack2 = append(stack2[:i2], stack2[i2+1:]...)
					stack1 = stack1[:len(stack1)-1]
					break
				}
			}
		}()
	}

	// we can only have one in each stack, otherwise the entry we received is invalid
	if (len(stack1) > 1) || (len(stack2) > 1) {
		return answer, fmt.Errorf("invalid entry")
	}

	// fill the answer with the lefting result from the popping pushing above
	answer = append(answer, stack1[0], stack2[0])

	return answer, nil
}
