package solutions

func Sol2(n int) int {
	sum := 2
	fib := []int{1, 2, 0}
	for fib[2] < n {
		fib[2] = fib[0] + fib[1]
		fib[0] = fib[1]
		fib[1] = fib[2]
		if fib[2]%2 == 0 {
			sum += fib[2]
		}
	}
	return sum
}
