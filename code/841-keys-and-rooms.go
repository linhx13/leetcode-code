package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	dfs(rooms, visited, 0)
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}

func dfs(rooms [][]int, visited []bool, cur int) {
	if visited[cur] {
		return
	}
	visited[cur] = true
	for _, v := range rooms[cur] {
		dfs(rooms, visited, v)
	}
}
