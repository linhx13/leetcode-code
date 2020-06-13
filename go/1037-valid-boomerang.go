package main

func isBoomerang(points [][]int) bool {
	if len(points) != 3 {
		return false
	}
	for i := 0; i < 3; i++ {
		if (points[i][0] == points[(i+1)%3][0]) && (points[i][1] == points[(i+1)%3][1]) {
			return false
		}
	}
	if (points[2][1]-points[0][1])*(points[1][0]-points[0][0])-(points[1][1]-points[0][1])*(points[2][0]-points[0][0]) == 0 {
		return false
	}
	return true
}
