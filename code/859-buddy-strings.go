package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		m := make(map[rune]bool)
		for _, r := range A {
			m[r] = true
		}
		return len(m) < len(A)
	}
	diff := []int{}
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
	}
	return len(diff) == 2 && A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]]
}
