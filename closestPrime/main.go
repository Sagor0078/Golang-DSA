

import (
	"fmt"
	"math"
)

func closestPrimes(left, right int) []int {
	if left > right {
		return []int{-1, -1}
	}

	sieve := make([]bool, right+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	for i := 2; i*i <= right; i++ {
		if sieve[i] {
			for j := i * i; j <= right; j += i {
				sieve[j] = false
			}
		}
	}

	primes := []int{}
	for i := left; i <= right; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	if len(primes) < 2 {
		return []int{-1, -1}
	}

	res := []int{-1, -1}
	minGap := math.MaxInt32

	for i := 1; i < len(primes); i++ {
		gap := primes[i] - primes[i-1]
		if gap < minGap {
			minGap = gap
			res = []int{primes[i-1], primes[i]}
		}
	}

	return res
}


