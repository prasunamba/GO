package module1

func isPowerOfThree(n int) bool {
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
