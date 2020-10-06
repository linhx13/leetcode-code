package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	cars := [][]int{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, []int{position[i], speed[i]})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})
	res := len(cars)
	i := len(cars) - 1
	for i > 0 {
		j := i - 1
		for ; j >= 0; j-- {
			if cars[j][1] < cars[i][1] {
				break
			}
			t1 := float64(cars[i][0]-cars[j][0]) / float64(cars[j][1]-cars[i][1])
			t2 := float64(target-cars[i][0]) / float64(cars[i][1])
			if t1 > t2 {
				break
			}
		}
		res -= i - (j + 1)
		i = j
	}
	return res
}

func main() {
	target := 10
	position := []int{0, 4, 2}
	speed := []int{2, 1, 3}
	// target := 12
	// position := []int{10, 8, 0, 5, 3}
	// speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(target, position, speed))
}
