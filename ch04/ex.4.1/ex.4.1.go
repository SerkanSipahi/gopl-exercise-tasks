package main

import (
	"crypto/sha256"
	"fmt"
)

// I have tried to resolve this exercise according ex4.1 (The Go Programming Language - Alan A. A. Donovan) but it was to complicated for me.
// I found a solution in https://github.com/SerkanSipahi/gopl-solutions/blob/master/ch04/ex4.2/ex4.2.go (without comments)
// and tries to understand what going on in init() and diffBits() see my comments

var pc [8]byte

func init() {

	// bitwise left shift(<<) operator
	// see: https://code.tutsplus.com/articles/understanding-bitwise-operators--active-11301

	// result of pc
	// byte(1 << 0) =   1 = 00000001
	// byte(1 << 1) =   2 = 00000010
	// byte(1 << 2) =   4 = 00000100
	// byte(1 << 3) =   8 = 00001000
	// byte(1 << 4) =  16 = 00010000
	// byte(1 << 5) =  32 = 00100000
	// byte(1 << 6) =  64 = 01000000
	// byte(1 << 7) = 128 = 10000000
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
}

func main() {

	hash1 := sha256.Sum256([]byte("hello world"))
	hash2 := sha256.Sum256([]byte("hello worlD"))
	printHash(hash1)
	printHash(hash2)
	diffBits(hash1, hash2)

}

func printHash(hash [32]byte) {

	for _, v := range hash {
		fmt.Printf("%X", v)
	}
	fmt.Println()

}

func diffBits(hash1, hash2 [32]byte) int {

	count := 0
	for i := 0; i < 32; i++ {
		byte1 := hash1[i] // e.g. = 154
		byte2 := hash2[i] // e.g. = 133

		for j := 0; j < 8; j++ {

			// @Note: between "byte" and "pc" is &(bitwise AND)
			// Good Explanation
			// see: https://code.tutsplus.com/articles/understanding-bitwise-operators--active-11301

			// #### When no bit diff exists ###
			// bitwise AND

			// ### byte1 ###
			// ================
			// deci  bit  = 128|64|32|16|8|4|2|1|
			// =========================================|
			// byte1      =  1 | 0| 1| 0|0|1|0|1| = 154
			// &pc[2]     =  0 | 0| 0| 0|0|1|0|0| =   4
			// =========================================|
			// bit diff   =    |  |  |  | |1| | | =   4

			// ### byte2 ###
			// deci  bit  = 128|64|32|16|8|4|2|1|
			// =========================================|
			// byte2      =  1 | 0| 0| 0|0|1|1|0| = 133
			// &pc[2]     =  0 | 0| 0| 0|0|1|0|0| =   4
			// =========================================|
			// bit diff   =    |  |  |  | |1| | |     4

			// 4 - 4 == 0 (no bit diff)

			/////////////////////////////////////////////
			/////////////////////////////////////////////

			// #### When bit diff exists ###

			// ### byte1 ###
			// ================
			// deci  bit  = 128|64|32|16|8|4|2|1|
			// =========================================|
			// byte1      =  1 | 0| 1| 0|0|1|0|1| = 154
			// &pc[5]     =  0 | 0| 1| 0|0|0|0|0| =  32
			// =========================================|
			// bit diff   =    |  | 1|  | | | | | =  32

			// ### byte2 ###
			// deci  bit  = 128|64|32|16|8|4|2|1|
			// =========================================|
			// byte2      =  1 | 0| 0| 0|0|1|1|0| = 133
			// &pc[5]     =  0 | 0| 1| 0|0|0|0|0| =  32
			// =========================================|
			// bit diff   =    |  | 0|  | | | | | =   0

			// 32 - 0 == 0 (has bit diff)

			if byte1&pc[j] != byte2&pc[j] {
				count++
			}
		}
	}
	return count
}
