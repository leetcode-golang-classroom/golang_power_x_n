package sol

func myPow(x float64, n int) float64 {
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	var Recursive func(x float64, n int) float64
	Recursive = func(x float64, n int) float64 {
		n = abs(n)
		if x == 0 {
			return 0
		}
		if x == 1 || n == 0 {
			return 1
		}
		res := Recursive(x, n/2)
		res = res * res
		if n%2 == 1 {
			res *= x
		}
		return res
	}
	result := Recursive(x, n)
	if n < 0 {
		return 1 / result
	}
	return result
}
