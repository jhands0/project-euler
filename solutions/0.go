package solutions

func Sol0(n int) int {
	sum := 0
	for i := 1; i < n+1; i += 2 {
		sum += i * i
	}
	return sum
}
