package main

func bitwiseComplement(N int) int {
    if N == 0 {
		return 1
	}
	res := 1
	for res -1 < N {
		res = res << 1
	}
	return res - N - 1
}
