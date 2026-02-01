package solutions

func Sol7(n int) int {
	limit := 1000000
	not_prime := make(map[int]bool)
	primes := []int{2}

	for i := 3; i <= limit; i += 2 {
		if _, ok := not_prime[i]; ok {
			continue
		}

		for j := i * 3; j <= limit; j += i * 2 {
			not_prime[j] = true
		}

		primes = append(primes, i)
	}
	return primes[n-1]
}
