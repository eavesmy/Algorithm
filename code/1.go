package code

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func Aplusb(a int, b int) int {
	// write your code here
	for b != 0 {
		res := (a & b) << 1
		a = a ^ b
		b = res
	}
	return a
}
