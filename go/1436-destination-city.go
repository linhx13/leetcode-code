package main

import "fmt"

func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for _, pair := range paths {
		m[pair[1]] = true
	}
	for _, pair := range paths {
		delete(m, pair[0])
	}
	for k, _ := range m {
		return k
	}
	return ""
}

func main() {
	paths := [][]string{}
	// paths = append(paths, []string{"London", "New York"})
	// paths = append(paths, []string{"New York", "Lima"})
	// paths = append(paths, []string{"Lima", "Sao Paulo"})
	// paths = append(paths, []string{"B", "C"})
	// paths = append(paths, []string{"D", "B"})
	// paths = append(paths, []string{"C", "A"})
	paths = append(paths, []string{"A", "Z"})
	fmt.Println(destCity(paths))
}
