package main

func binaryGap(N int) int {
	max := 0
	last := -1
	for i := 0; N > 0; N, i = N>>1, i+1 {
		if N&1 == 1 {
			if last != -1 {
				d := i - last
				if d > max {
					max = d
				}
			}
			last = i
		}
	}
	return max
}
