package main

import "fmt"

// 寻找数组中心索引，数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
func pivotIndex(nums []int) int {
	sum := 0
	size := len(nums)
	for i := 0; i < size; i++ {
		sum += nums[i]
	}

	halfsum := 0
	for i := 0; i < size; i++ {
		if (halfsum*2 + nums[i]) == sum {
			return i
		}
		halfsum += nums[i]
	}

	return -1
}

func main() {
	//fmt.Println("hello world")

	nums := []int{1, 3, 5}
	fmt.Println(pivotIndex(nums))
}
