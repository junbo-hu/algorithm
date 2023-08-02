package main

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

import "fmt"

// 低位开始计算，如果是9，则赋值0并计算高位；如果是0~8，则+1，并结束计算，返回数组
func plusOne(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nums
	}
	for i := size - 1; i >= 0; i-- {
		if nums[i] == 9 {
			nums[i] = 0

			if i == 0 {
				nums = append([]int{1}, nums...)
				return nums
			}
		} else {
			nums[i]++
			return nums
		}
	}

	return nums
}

func main() {
	nums := []int{0}
	result := plusOne(nums)
	fmt.Println(result)
}
