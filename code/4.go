package code

import "fmt"

/**
 * @param n: An integer
 * @return: return a  integer as description.
 */
func NthUglyNumber(n int) int {
	// write your code here

	// 被 5 3 2 整除剩 1

	// 准备一个排好序的数组,直接取就可以，空间换时间

	if n == 1 {
		return 1
	}

	ret := make([]int, n)
	for i := 1; i < n; i++ {
		ret[i-1] = i
	}

	a := 0
	b := 0
	c := 0
	min := 0

	for i := 1; i < n; i++ {

		tmpA := ret[a] * 2
		tmpB := ret[b] * 3
		tmpC := ret[c] * 5

		min = tmpC

		fmt.Println(min, tmpA, tmpB, tmpC)

		if tmpB < tmpC {
			min = tmpB
		}
		if tmpA < min {
			min = tmpA
		}

		if min == tmpA {
			a++
		}
		if min == tmpB {
			b++
		}
		if min == tmpC {
			c++
		}

		ret[i] = min
	}

	return ret[n-1]
}
