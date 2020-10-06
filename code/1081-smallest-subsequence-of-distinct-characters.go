package main

func smallestSubsequence(text string) string {
	count := make([]int, 26)
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	stack := []byte{}
	for i := 0; i < len(text); i++ {
		c := text[i]
		count[c-'a']--
		hit := false
		for _, x := range stack {
			if x == c {
				hit = true
				break
			}
		}
		if hit {
			continue
		}
		if len(stack) > 0 {
			for p := stack[len(stack)-1]; c < p && count[p-'a'] > 0; p = stack[len(stack)-1] {
				stack = stack[0 : len(stack)-1]
				if len(stack) == 0 {
					break
				}
			}
		}
		stack = append(stack, c)
	}
	return string(stack)
}
