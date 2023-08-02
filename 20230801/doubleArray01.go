package main

import (
	"fmt"
)

func dominantIndex(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])
	if cols == 0 {
		return nil
	}

	out := []int{}
	//初始方向
	horizon := 1
	sum := 0
	pos := 0 //方向正，则表示x，方向负，则表示y

	for i := 0; i < rows+cols-1; i++ {
		if horizon == 1 {
			//正方向的时候，pos是x坐标，起点在左边界或下边界
			if sum < rows {
				pos = 0
			} else {
				pos = sum - rows + 1
			} //确认起点
			//遍历从起点开始直到边界
			for pos < cols && sum-pos > -1 {
				out = append(out, matrix[sum-pos][pos])
				pos++
			}
			//遍历完了，转向并进入下一条对角线
			sum++
			horizon *= -1
		} else {
			//反方向的时候，pos是y坐标，起点在右边界或上边界
			if sum < cols {
				pos = 0
			} else {
				pos = sum - cols + 1
			} //确认起点
			//遍历从起点开始直到边界
			for pos < rows && sum-pos > -1 {
				out = append(out, matrix[pos][sum-pos])
				pos++
			}
			//遍历完了，转向并进入下一条对角线
			sum++
			horizon *= -1
		}
	}
	return out
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	result := dominantIndex(matrix)
	fmt.Println(result)
}
