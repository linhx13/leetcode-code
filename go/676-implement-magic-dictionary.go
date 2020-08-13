package main

type MagicDictionary struct {
	s map[string]bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{s: make(map[string]bool)}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		this.s[word] = true
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	chars := []rune(word)
	for i, t := range chars {
		for c := 'a'; c <= 'z'; c++ {
			if c == t {
				continue
			}
			chars[i] = c
			v := this.s[string(chars)]
			if v {
				return true
			}
		}
		chars[i] = t
	}
	return false
}
