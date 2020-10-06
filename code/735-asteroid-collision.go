package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, x := range asteroids {
		if x > 0 {
			stack = append(stack, x)
		} else if x < 0 {
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, x)
			} else {
				valid := true
				for len(stack) > 0 && stack[len(stack)-1] > 0 && valid {
					if -1*x > stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
					} else if -1*x == stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
						valid = false
					} else {
						valid = false
					}
				}
				if valid {
					stack = append(stack, x)
				}
			}
		}
	}
	return stack
}

func main() {
	// nums := []int{5, 10, -5}
	// nums := []int{8, -8}
	// nums := []int{10, 2, -5}
	// nums := []int{-2, -1, 1, 2}
	nums := []int{}
	fmt.Println(asteroidCollision(nums))
}
