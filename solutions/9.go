package solutions

import "math"

func Sol9(n int) int {
	upper_limit := 3000
	cache := make(map[int]int)

	for a := 1; a < upper_limit; a++ {
		for b := a + 1; b < upper_limit-a; b++ {

			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2)))
			if c*c != c2 {
				continue
			}

			sum := a + b + c
			if sum > upper_limit {
				break
			}

			if cache[sum] < a*b*c {
				cache[sum] = a * b * c
			}
		}
	}

	return cache[n]
}
