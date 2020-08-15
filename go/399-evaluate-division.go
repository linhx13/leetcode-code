package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		x := equations[i][0]
		y := equations[i][1]
		if _, ok := g[x]; !ok {
			g[x] = make(map[string]float64)
		}
		g[x][y] = values[i]
		if _, ok := g[y]; !ok {
			g[y] = make(map[string]float64)
		}
		g[y][x] = 1.0 / values[i]
	}

	res := []float64{}
	for _, q := range queries {
		x, y := q[0], q[1]
		_, ok_x := g[x]
		_, ok_y := g[y]
		if !ok_x || !ok_y {
			res = append(res, -1.0)
			continue
		}
		v := make(map[string]bool)
		res = append(res, dfs(x, y, g, v))
	}
	return res
}

func dfs(x string, y string, g map[string]map[string]float64, v map[string]bool) float64 {
	if x == y {
		return 1.0
	}
	v[x] = true
	for k, _ := range g[x] {
		if v[k] {
			continue
		}
		d := dfs(k, y, g, v)
		if d > 0 {
			return d * g[x][k]
		}
	}
	return -1.0
}
