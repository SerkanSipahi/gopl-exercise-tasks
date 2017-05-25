package main

import (
	"bytes"
	"fmt"
)

func main() {

	word1 := []byte("serkan")
	word2 := []byte("nakres")

	fmt.Println(IsAnagram(word1, word2))
}

func IsAnagram(word1 []byte, word2 []byte) bool {

	m := make(map[int]int)

	for _, y := range bytes.ToLower(word2) {
		m[int(y)]++
	}

	var ret []int
	for _, x := range bytes.ToLower(word1) {
		if m[int(x)] > 0 {
			m[int(x)]--
			continue
		}
		ret = append(ret, int(x))
	}

	if len(ret) == 0 {
		return true
	} else {
		return false
	}
}
