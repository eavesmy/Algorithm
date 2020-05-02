package code

// import "fmt"

/**
 * @param prices: Given an integer array
 * @return: Maximum profit
 */
func MaxProfit(prices []int) int {
	// write your code here

	// 两层遍历

	buy := -1
	sell := -1

	dis := 0

	// sell 下标比 buy >=

	for i, item := range prices {
		buy = item

		for j := i; j < len(prices); j++ {
			sell = prices[j]

			if (sell - buy) > dis {
				dis = sell - buy
			}
		}
	}

	return dis
}
