package main

import "fmt"

func numTilePossibilities(tiles string) int {
	res := make(map[string]bool)
	chars := make(map[byte]int)
	for i := 0; i < len(tiles); i++ {
		chars[tiles[i]]++
	}
	helper(chars, &[]byte{}, res)
	return len(res)
}

func helper(chars map[byte]int, buf *[]byte, res map[string]bool) {
	if len(*buf) > 0 {
		tmp := make([]byte, len(*buf))
		copy(tmp, *buf)
		res[string(tmp)] = true
	}
	flag := false
	for c, n := range chars {
		if n == 0 {
			continue
		}
		flag = true
		chars[c]--
		*buf = append(*buf, c)
		helper(chars, buf, res)
		*buf = (*buf)[0 : len(*buf)-1]
		chars[c]++
	}
	if !flag {
		tmp := make([]byte, len(*buf))
		copy(tmp, *buf)
		res[string(tmp)] = true
	}
}

func main() {
	// tiles := "AAB"
	tiles := "AAABBC"
	// tiles := "V"
	fmt.Println(numTilePossibilities(tiles))
}
