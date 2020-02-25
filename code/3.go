package code

/**
 * @param k: An integer
 * @param n: An integer
 * @return: An integer denote the count of digit k in 1..n
 */
func DigitCounts(k int, n int) int {
	// 利用 % 取余
	count := 0
	for i := 0; i <= n; i++ {
		j := i
		for {
			if j%10 == k {
				count++
			}

			j /= 10

			if j == 0 {
				break // 除完了
			}
		}
	}

	return count
}
