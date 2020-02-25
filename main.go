package main

import "fmt"
import code "./code"

func main() {
	/*
		ret64 := code.TrailingZeros(5)
		fmt.Println(ret64)
	*/
	// ret_arr := GetNarcissisticNumbers(1)

	ret := code.DigitCounts(1, 15) // 1 10 11 11 12 13 14 15
	fmt.Println(ret)
}
