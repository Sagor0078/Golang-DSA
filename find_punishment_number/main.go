package main

import (
	"fmt"
	"strconv"
)

func rec(i, rem, sm, num int, s string) bool {
	if i == len(s) {
		return sm+rem == num
	}

	digit := int(s[i] - '0')
	newRem := rem*10 + digit

	return rec(i+1, 0, sm+newRem, num, s) || rec(i+1, newRem, sm, num, s)
}

func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		num := i
		s := strconv.Itoa(i * i)

		if rec(0, 0, 0, num, s) {
			ans += i * i
		}
	}
	return ans
}


