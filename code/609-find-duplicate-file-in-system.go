package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {

	m := make(map[string][]string)
	for _, path := range paths {
		arr := strings.Fields(path)
		dir := arr[0]
		for _, item := range arr[1:len(arr)] {
			brr := strings.Split(item, "(")
			filename := brr[0]
			content := strings.Split(brr[1], ")")[0]
			m[content] = append(m[content], dir+"/"+filename)
		}
	}
	res := [][]string{}
	for _, v := range m {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	fmt.Println(findDuplicate(paths))
}
