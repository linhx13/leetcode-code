package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := []string{}
	res = append(res, folder[0])
	last := strings.Split(folder[0], "/")
	for i := 1; i < len(folder); i++ {
		cur := strings.Split(folder[i], "/")
		if isSub(last, cur) {
			continue
		}
		res = append(res, folder[i])
		last = cur
	}
	return res
}

func isSub(x []string, y []string) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	// folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	// folder := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	// folder := []string{"/a", "/a/b/c", "/a/b/d"}
	folder := []string{"/a"}
	fmt.Println(removeSubfolders(folder))
}
