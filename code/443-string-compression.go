package main

func appendChars(chars []byte, last byte, cnt int, res int) int {
	chars[res] = last
	res++
	if cnt == 1 {
		return res
	}
	r := res
	for cnt > 0 {
		chars[res] = byte('0' + (cnt % 10))
		res++
		cnt = cnt / 10
	}
	for i := 0; i < (res-r)/2; i++ {
		chars[r+i], chars[res-i-1] = chars[res-i-1], chars[r+i]
	}
	return res
}

func compress(chars []byte) int {
	var last byte
	cnt := 0
	res := 0
	for _, char := range chars {
		if last != byte(0) && last != char {
			res = appendChars(chars, last, cnt, res)
			last = byte(0)
			cnt = 0
		}
		last = char
		cnt++
	}
	if last != byte(0) {
		res = appendChars(chars, last, cnt, res)
	}
	return res
}
