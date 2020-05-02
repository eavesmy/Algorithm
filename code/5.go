package code

/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 */

// 基本思想还是排序，完用n做下标去取。大家都用快排做，我用选择排序做个。

func KthLargestElement(n int, nums []int) int {
	length := len(nums)
	max := -1
	for i := 0; i < length; i++ {
		max = i
		for j := i + 1; j < length; j++ {
			if nums[j] > nums[max] {
				max = j
			}
		}
		nums[i], nums[max] = nums[max], nums[i]

		if i >= n {
			return nums[n-1]
		}
	}
	return nums[n-1]
}
