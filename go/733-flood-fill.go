package main

import "fmt"

func dfs(image [][]int, sr int, sc int, newColor int, oldColor int, visited [][]bool) {
	if visited[sr][sc] {
		return
	}
	visited[sr][sc] = true
	if image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	if sr-1 >= 0 {
		dfs(image, sr-1, sc, newColor, oldColor, visited)
	}
	if sr+1 < len(image) {
		dfs(image, sr+1, sc, newColor, oldColor, visited)
	}
	if sc-1 >= 0 {
		dfs(image, sr, sc-1, newColor, oldColor, visited)
	}
	if sc+1 < len(image[0]) {
		dfs(image, sr, sc+1, newColor, oldColor, visited)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := [][]bool{}
	for i := 0; i < len(image); i++ {
		visited = append(visited, make([]bool, len(image[0])))
	}
	dfs(image, sr, sc, newColor, image[sr][sc], visited)
	return image
}

func main() {
	image := [][]int{}
	// image = append(image, []int{1, 1, 1})
	// image = append(image, []int{1, 1, 0})
	// image = append(image, []int{1, 0, 1})
	image = append(image, []int{1})
	image = append(image, []int{0})
	image = append(image, []int{1})
	fmt.Println(floodFill(image, 0, 0, 2))
}
