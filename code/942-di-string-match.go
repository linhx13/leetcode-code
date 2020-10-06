package main

func diStringMatch(S string) []int {
	res := []int{}
	min, max := 0, len(S)
	for _, c := range S {
		if c == 'I' {
			res = append(res, min)
			min++
		} else {
			res = append(res, max)
			max--
		}
	}
	if S[len(S)-1] == 'D' {
		res = append(res, res[len(res)-1]-1)
	} else {
		res = append(res, res[len(res)-1]+1)
	}
	return res
}
