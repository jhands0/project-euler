package solutions

func checkDivisible(i int, div []int) bool {
	for _, m := range div {
		if i%m != 0 {
			return false
		}
	}
	return true
}

func Sol5(n int) int {
	divisors := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	i := 2520
	for !checkDivisible(i, divisors) {
		i += 2520
	}
	return i
}
