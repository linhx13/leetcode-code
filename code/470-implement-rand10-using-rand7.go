package main

func rand10() int {
	c := (rand7()-1)*7 + rand7() - 1
	if c >= 40 {
		return rand10()
	} else {
		return (c % 10) + 1
	}
}
