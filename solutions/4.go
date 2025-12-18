package solutions

import "strconv"

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Sol4(n int) int {
	largest := 0
	for i := 100; i < 1000; i++ {
		for j := 110; j < 1000; j += 11 {
			num := i * j
			str := strconv.Itoa(num)
			if str == reverse(str) {
				if largest < num {
					largest = num
				}
			}
		}
	}
	return largest
}
