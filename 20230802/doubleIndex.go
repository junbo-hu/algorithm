package main

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	var i, j, res int
	i = 0
	j = len(height) - 1
	res = 0
	for i < j {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
		} else {
			res = max(res, height[j]*(j-i))
			j--
		}
	}

	return res
}

func main() {
	array := []int{10, 8, 6, 2, 1, 3, 5, 7, 4}

	fmt.Println(maxArea(array))
}
