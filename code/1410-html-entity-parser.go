package main

import "fmt"

func entityParser(text string) string {
	m := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}

	res := make([]rune, 0, len(text))
	last := -1
	for _, c := range text {
		res = append(res, c)
		if c == '&' {
			last = len(res) - 1
		}
		if c != ';' {
			continue
		}
		if last != -1 {
			v, ok := m[string(res[last:])]
			if ok {
				res = res[:last]
				for _, c := range v {
					res = append(res, c)
				}
			}
			last = -1
		}
	}
	return string(res)
}

func main() {
	// text := "&amp; is an HTML entity but &ambassador; is not."
	// text := "and I quote: &quot;...&quot;"
	// text := "Stay home! Practice on Leetcode :)"
	// text := "x &gt; y &amp;&amp; x &lt; y is always false"
	// text := "leetcode.com&frasl;problemset&frasl;all"
	text := "&amp;amp;"
	fmt.Println(entityParser(text))
}
