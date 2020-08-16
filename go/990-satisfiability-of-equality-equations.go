package main

import "fmt"

func equationsPossible(equations []string) bool {
	g := make(map[string]map[string]int)
	if !buildGraph(equations, g) {
		return false
	}
	for _, e := range equations {
		x := e[0:1]
		y := e[3:4]
		rel := e[1:3]
		v := make(map[string]bool)
		d := dfs(x, y, g, v)
		if rel == "==" && d < 0 {
			return false
		}
		if rel == "!=" && d > 0 {
			return false
		}
	}
	return true
}

func buildGraph(equations []string, g map[string]map[string]int) bool {
	for i := 0; i < len(equations); i++ {
		x := equations[i][0:1]
		y := equations[i][3:4]
		rel := 1
		if equations[i][1:3] == "!=" {
			rel = -1
		}
		if _, ok := g[x]; !ok {
			g[x] = make(map[string]int)
		}
		v := g[x][y]
		if v != 0 {
			if v != rel {
				return false
			}
		} else {
			g[x][y] = rel
		}
		if _, ok := g[y]; !ok {
			g[y] = make(map[string]int)
		}
		v = g[y][x]
		if v != 0 {
			if v != rel {
				return false
			}
		} else {
			g[y][x] = rel
		}
	}
	return true
}

func dfs(x string, y string, g map[string]map[string]int, v map[string]bool) int {
	if x == y {
		return 1
	}
	v[x] = true
	for k, t := range g[x] {
		if v[k] {
			continue
		}
		if t == -1 {
			continue
		}
		d := dfs(k, y, g, v)
		if d > 0 {
			return d
		}
	}
	return -1
}

func main() {
	// equations := []string{"b==a", "a==b"}
	// equations := []string{"a==b", "b!=a"}
	// equations := []string{"a==b", "b==c", "a==c"}
	// equations := []string{"a==b", "b!=c", "c==a"}
	equations := []string{"c==c", "b==d", "x!=z"}
	fmt.Println(equationsPossible(equations))
}
