package main

type State int

const (
	UNKNOWN = iota
	VISITING
	SAFE
	UNSAFE
)

func eventualSafeNodes(graph [][]int) []int {
	states := make([]State, len(graph))
	res := []int{}
	for i := 0; i < len(graph); i++ {
		if dfs(graph, i, states) == SAFE {
			res = append(res, i)
		}
	}
	return res
}

func dfs(graph [][]int, cur int, states []State) State {
	if states[cur] == VISITING {
		states[cur] = UNSAFE
		return states[cur]
	}
	if states[cur] != UNKNOWN {
		return states[cur]
	}
	states[cur] = VISITING
	for _, next := range graph[cur] {
		if dfs(graph, next, states) == UNSAFE {
			states[cur] = UNSAFE
			return states[cur]
		}
	}
	states[cur] = SAFE
	return states[cur]
}
