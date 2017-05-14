package main

import (
	"bytes"
	"strings"
)

// Solution over stackoverflow but extended for exercise:
// http://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma#26939091

func comma(value string, sep ...rune) string {

	_sep := ','
	if len(sep) == 1 {
		_sep = sep[0]
	}

	l := len(value)

	var commaIndex int
	floatPoint := strings.Index(value, ".")
	if floatPoint >= 0 {
		commaIndex = 3 - (floatPoint % 3)
	} else {
		commaIndex = 3 - (l % 3)
	}

	if commaIndex == 3 {
		commaIndex = 0
	}

	var buff bytes.Buffer
	for i := 0; i < l; i++ {
		if commaIndex == 3 && i < floatPoint {
			buff.WriteRune(_sep)
			commaIndex = 0
		}
		commaIndex++
		buff.WriteByte(value[i])
	}
	return buff.String()
}

func main() {
	println(comma("45638522.34232874", ',')) // 45,638,522.34232874
	println(comma("4563852234232.874", ',')) // 4,563,852,234,232.874
}
