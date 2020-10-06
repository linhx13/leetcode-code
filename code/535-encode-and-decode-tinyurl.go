package main

import (
	"math/rand"
	"strings"
)

type Codec struct {
	short2long map[string]string
	long2short map[string]string
	vocab      []byte
}

func Constructor() Codec {
	return Codec{
		long2short: make(map[string]string),
		short2long: make(map[string]string),
		vocab:      []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if v, ok := this.long2short[longUrl]; ok {
		return v
	}
	randBytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		randBytes = append(randBytes, this.vocab[rand.Intn(len(this.vocab))])
	}
	if _, ok := this.short2long[string(randBytes)]; ok {
		randBytes[rand.Intn(6)] = this.vocab[rand.Intn(len(this.vocab))]
	}
	randStr := string(randBytes)
	this.short2long[randStr] = longUrl
	this.long2short[longUrl] = randStr
	return "http://tinyurl.com/" + randStr
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	parts := strings.Split(shortUrl, "/")
	encoded := parts[len(parts)-1]
	if v, ok := this.short2long[encoded]; ok {
		return v
	} else {
		return shortUrl
	}
}
