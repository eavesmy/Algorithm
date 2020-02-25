package code

// import "fmt"

/**
 * @param n: A long integer
 * @return: An integer, denote the number of trailing zeros in n!
 */
func TrailingZeros(n int64) int64 {

	var ret int64 = 0

	for n != 0 {
		n /= 5
		ret += n
	}

	// 直接判断能被几个 5 整除

	return ret
}
