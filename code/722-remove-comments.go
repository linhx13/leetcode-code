package main

func removeComments(source []string) []string {
	inBlock := false
	res := []string{}
	var buf []byte
	for _, line := range source {
		i := 0
		if !inBlock {
			buf = []byte{}
		}
		for i < len(line) {
			if !inBlock && i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
				inBlock = true
				i++
			} else if inBlock && i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
				inBlock = false
				i++
			} else if !inBlock && i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
				break
			} else if !inBlock {
				buf = append(buf, line[i])
			}
			i++
		}
		if !inBlock && len(buf) > 0 {
			res = append(res, string(buf))
		}
	}
	return res
}
