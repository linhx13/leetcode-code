package main

import (
	"fmt"
	"strings"
)

func maskPII(S string) string {
	if strings.Index(S, "@") != -1 {
		// email
		arr := strings.Split(S, "@")
		name1 := strings.ToLower(arr[0])
		name2 := strings.ToLower(arr[1])
		res := []byte{}
		res = append(res, name1[0])
		res = append(res, []byte("*****")...)
		res = append(res, name1[len(name1)-1])
		res = append(res, '@')
		res = append(res, name2...)
		return string(res)
	} else {
		// phone number
		for _, k := range []string{"+", "-", "(", ")", " "} {
			S = strings.ReplaceAll(S, k, "")
		}
		res := []byte{}

		if len(S) > 10 {
			res = append(res, '+')
		}
		for i := 0; i < len(S)-10; i++ {
			res = append(res, '*')
		}
		if len(S) > 10 {
			res = append(res, '-')
		}
		res = append(res, []byte("***-***-")...)
		res = append(res, S[len(S)-4:len(S)]...)
		return string(res)
	}
}

func main() {
	// S := "LeetCode@LeetCode.com"
	// S := "AB@qq.com"
	// S := "1(234)567-890"
	S := "+86-(10)12345678"
	fmt.Println(maskPII(S))
}
