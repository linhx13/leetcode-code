package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := []rune("0123456789abcdef")
	n := int64(num)
	if num < 0 {
		n += int64(1) << 32
	}
	res := []rune{}
	for n > 0 {
		res = append([]rune{hex[n%16]}, res...)
		n = n / 16
	}
	return string(res)
}
