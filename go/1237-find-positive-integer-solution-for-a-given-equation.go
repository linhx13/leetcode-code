package main

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := [][]int{}
	x, y := 1, 1000
	for x <= 1000 && y >= 1 {
		r := customFunction(x, y)
		if r > z {
			y--
		} else if r < z {
			x++
		} else {
			res = append(res, []int{x, y})
			x++
		}
	}
	return res
}
