package sms

import (
	"fmt"
	"unicode/utf8"
)

func NigeriaNumber(number string) string {
	if len(number) != 11 {
		fmt.Println("In correct number.")
		return number
	}
	newNumber := trimFirstRune(number)
	return "234" + newNumber
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}


/*func TrimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}*/

