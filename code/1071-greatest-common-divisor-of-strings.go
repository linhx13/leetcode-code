package main

func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) == 0 {
		return str1
	}
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	for i := 0; i < len(str2); i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}
	return gcdOfStrings(str1[len(str2):], str2)
}
