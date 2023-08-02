package main

import "fmt"

func yangHui(num int) [][]int {
	var out [][]int
	if num <= 0 {
		return out
	}
	for row := 0; row < num; row++ {
		rowArray := make([]int, row+1)
		for col := 0; col <= row; col++ {
			if col == 0 || row == col {
				rowArray[col] = 1
			} else {
				rowArray[col] = out[row-1][col] + out[row-1][col-1]
			}
		}
		out = append(out, rowArray)
	}

	return out
}

func main() {
	num := 5
	result := yangHui(num)

	fmt.Println(result)
}
