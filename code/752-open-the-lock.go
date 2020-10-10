package main

import "fmt"

func openLock(deadends []string, target string) int {
	dm := make(map[string]bool)
	for _, d := range deadends {
		dm[d] = true
	}
	visited := make(map[string]bool)
	if dm["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	queue := []string{}
	queue = append(queue, "0000")
	visited["0000"] = true
	res := 0
	for len(queue) > 0 {
		for sz := len(queue); sz > 0; sz-- {
			cur := queue[0]
			queue = queue[1:len(queue)]
			if cur == target {
				return res
			}
			for i := 0; i < 4; i++ {
				up := []byte(cur)
				if up[i] == '9' {
					up[i] = '0'
				} else {
					up[i]++
				}
				upStr := string(up)
				if !visited[upStr] && !dm[upStr] {
					queue = append(queue, upStr)
					visited[upStr] = true
				}
				down := []byte(cur)
				if down[i] == '0' {
					down[i] = '9'
				} else {
					down[i]--
				}
				downStr := string(down)
				if !visited[downStr] && !dm[downStr] {
					queue = append(queue, downStr)
					visited[downStr] = true
				}
			}
		}
		res++
	}
	return -1
}

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
