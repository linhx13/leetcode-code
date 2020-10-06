package main

import "fmt"

func getFolderNames(names []string) []string {
	m := make(map[string]int)
	res := []string{}
	for _, name := range names {
		if v, ok := m[name]; !ok {
			res = append(res, name)
			m[name] = 1
		} else {
			newName := fmt.Sprintf("%s(%d)", name, v)
			for _, ok := m[newName]; ok; {
				v++
				newName = fmt.Sprintf("%s(%d)", name, v)
				_, ok = m[newName]
			}
			m[name] = v + 1
			m[newName] = 1
			res = append(res, newName)
		}
	}
	return res
}

func main() {
	// names := []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}
	// names := []string{"wano", "wano", "wano", "wano"}
	names := []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
	fmt.Println(getFolderNames(names))
}
