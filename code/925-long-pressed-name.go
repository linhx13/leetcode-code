package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 && len(typed) != 0 {
		return false
	}
	if len(name) != 0 && len(typed) == 0 {
		return false
	}
	runes_name := []rune(name)
	runes_typed := []rune(typed)
	char_name, cnt_name := '-', 0
	char_typed, cnt_typed := '-', 0
	i, j := 0, 0
	for i < len(runes_name) && j < len(runes_typed) {
		for i < len(runes_name) {
			if runes_name[i] != char_name && cnt_name != 0 {
				break
			}
			char_name = runes_name[i]
			cnt_name++
			i++
		}
		for j < len(runes_typed) {
			if runes_typed[j] != char_typed && cnt_typed != 0 {
				break
			}
			char_typed = runes_typed[j]
			cnt_typed++
			j++
		}
		if !(char_name == char_typed && cnt_name <= cnt_typed) {
			return false
		}
		char_name, cnt_name = '-', 0
		char_typed, cnt_typed = '-', 0
	}
	if (i >= len(runes_name) && j < len(runes_typed)) || (i < len(runes_name) && j >= len(runes_typed)) {
		return false
	}
	return true
}

func main() {
	name := ""
	typed := ""
	fmt.Println(isLongPressedName(name, typed))
}
