package main

ifunc numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		m := make(map[int]int)
		for j := 0; j < len(points); j++ {
			a := points[i][0] - points[j][0]
			b := points[i][1] - points[j][1]
			m[a*a+b*b]++
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}
	return res
}
