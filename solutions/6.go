package solutions

func Sol6(n int) int {
	sum_of_squares := 0
	sum := 0
	for i := 1; i < n+1; i++ {
		sum_of_squares += i * i
		sum += i
	}
	square_of_sum := sum * sum

	return square_of_sum - sum_of_squares
}
