package main

import (
	"unicode"
)

func clearDigits(s string) string {
	res := ""
	for _, x := range s {
		if unicode.IsDigit(x) {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res += string(x)
		}
	}
	return res
}
