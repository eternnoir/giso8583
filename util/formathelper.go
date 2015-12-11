package util

import (
	"strings"
)

// RightPad help pad some string to origin string's right repeat until total length.
func PadRigtht(s string, padStr string, totalLen int) string {
	var padCountInt int
	padCountInt = 1 + ((totalLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:totalLen]
}

// LeftPad help pad some string to origin string's left repeat until total length.
func PadLeft(s string, padStr string, totalLen int) string {
	var padCountInt int
	padCountInt = 1 + ((totalLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - totalLen):]
}
