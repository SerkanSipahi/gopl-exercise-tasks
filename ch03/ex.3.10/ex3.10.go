package main

import (
	"bytes"
)

// Solution:
// http://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma#26939091
// @Fixme: add tests and benchmark tests

func comma(s string) string {

	startOffset := 0
	var buff bytes.Buffer
	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)
	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {
		if commaIndex == 3 {
			buff.WriteRune(',')
			commaIndex = 0
		}
		commaIndex++
		buff.WriteByte(s[i])
	}
	return buff.String()
}

func main() {
	println(comma("45638522874")) // 45,638,522,874
	println(comma("5638522874"))  //  5,638,522,874
}
