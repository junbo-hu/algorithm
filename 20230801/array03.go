package main

import "fmt"

func findMax(array []int) int {
	max := -1
	secondMax := -1
	index := 0
	if len(array) < 2 {
		return 0
	}
	// 找出array中的最大值和次大值
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			secondMax = max
			max = array[i]
			index = i
		} else if array[i] > secondMax {
			secondMax = array[i]
		}
	}

	// 最大值是次大值的两倍以上
	if max >= secondMax*2 {
		return index
	}
	return -1
}

func main() {
	array := []int{9, 18, 100, 34}

	result := findMax(array)

	fmt.Println(result)
}
