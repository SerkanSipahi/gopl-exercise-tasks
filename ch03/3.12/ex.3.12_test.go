package main

import (
	"testing"
)

type testpair struct {
	word1     []byte
	word2     []byte
	isAnagram bool
}

var tests = []testpair{
	{[]byte("serkan"), []byte("nakres"), true},
	{[]byte("liya"), []byte("ayil"), true},
	{[]byte("abc"), []byte("cda"), false},
	{[]byte("ABC"), []byte("cba"), true},
}

func TestIsAnagram(t *testing.T) {
	for _, pair := range tests {
		value := IsAnagram(pair.word1, pair.word2)
		if value == pair.isAnagram {
			continue
		}

		t.Error(
			"For", pair.word1, pair.word2,
			"expected", pair.isAnagram,
			"got", value,
		)
	}
}
