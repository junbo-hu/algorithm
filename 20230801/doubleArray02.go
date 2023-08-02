package main

import "fmt"

// 顺时针螺旋遍历二维数组
func spiralOrder(array [][]int) []int {
	var out []int
	rows := len(array)
	if rows == 0 {
		return out
	}
	cols := len(array[0])
	if cols == 0 {
		return out
	}

	x := 0
	y := 0
	xMin := 0
	xMax := cols - 1
	yMin := 0
	yMax := rows - 1
	xOff := 1
	yOff := 0

	for i := 0; i < rows*cols; i++ {
		out = append(out, array[y][x])
		x += xOff
		y += yOff

		//到右边界
		if x > xMax {
			x = xMax
			yOff = 1
			xOff = 0
			y++
			yMin++
		}

		//到下边界
		if y > yMax {
			y = yMax
			x--
			xOff = -1
			yOff = 0
			xMax--
		}
		//到左边界
		if x < xMin {
			x = xMin
			y--
			xOff = 0
			yOff = -1
			yMax--
		}

		//到上边界
		if y < yMin {
			y = yMin
			x++
			yOff = 0
			xOff = 1
			xMin++
		}

	}

	return out
}

func main() {
	array := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	result := spiralOrder(array)

	fmt.Println(result)
}
