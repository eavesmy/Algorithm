package main

import "fmt"
import "time"
import code "./code"

func main() {
	start := time.Now()
	/*
					ret64 := code.TrailingZeros(5)
					fmt.Println(ret64)

					// ret_arr := GetNarcissisticNumbers(1)

					ret := code.DigitCounts(1, 15) // 1 10 11 11 12 13 14 15
					fmt.Println(ret)

				ret := code.NthUglyNumber(2)
				fmt.Println("ret", ret)

			ret := code.KthLargestElement(2, []int{9, 3, 2, 4, 8, 10}) // 3
			fmt.Println("ret", ret)

		ret := code.MaxProfit([]int{3, 2, 3, 1, 2})
		fmt.Println(ret)
	*/

	ret := code.KthLargestElement(3, []int{9, 3, 2, 4, 8}) // 4

	fmt.Println(ret)

	end := time.Now()

	unit := "纳秒"
	dis := end.UnixNano() - start.UnixNano()

	if dis > 100000 {
		unit = "微秒"
		dis /= 100000
	}

	fmt.Println("耗时：", dis, unit)
}
