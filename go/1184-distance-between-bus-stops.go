package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	dist1 := 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		dist1 += distance[i]
	}
	dist2 := 0
	for i := start; i != destination; i = (i - 1 + len(distance)) % len(distance) {
		dist2 += distance[(i-1+len(distance))%len(distance)]
	}
	if dist1 < dist2 {
		return dist1
	} else {
		return dist2
	}
}
