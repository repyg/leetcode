func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	t1, t2, t3 := 0, 1, 1
	for i := 3; i < n; i++ {
		t1, t2, t3 = t2, t3, t1+t2+t3
	}
	return t1 + t2 + t3
}