// Package main for matrix question
package main

import "fmt"

var island = [][]int{
	{1, 0, 0, 0, 0, 1}, // 0
	{0, 1, 0, 1, 1, 1}, // 1
	{0, 0, 1, 0, 1, 0}, // 2
	{1, 1, 0, 0, 1, 0}, // 3
	{1, 0, 1, 1, 0, 0}, // 4
	{1, 0, 0, 0, 0, 1}, // 2

}

func main() {
	removeIslands(island)
	printMap(island)
}

// Remove the islands
func removeIslands(arr [][]int) {
	n, m := len(arr), len(arr[0])
	// Iterate through the array
	for i := 0; i < m; i++ {
		if arr[0][i] == 1 {
			visitAndMark(arr, 0, i)
		}
		if arr[n-1][i] == 1 {
			visitAndMark(arr, n-1, i)
		}
	}

	for i := 0; i < n; i++ {
		if arr[i][0] == 1 {
			visitAndMark(arr, i, 0)
		}
		if arr[i][m-1] == 1 {
			visitAndMark(arr, i, m-1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 2 {
				arr[i][j] = 1
			} else if arr[i][j] == 1 {
				arr[i][j] = 0
			}
		}
	}
}

func printMap(arr [][]int) {
	for _, row := range arr {
		for _, col := range row {
			fmt.Printf("[%v]", col)
		}
		fmt.Println()
	}
}

// visit and mark the borders
func visitAndMark(arr [][]int, x int, y int) {
	n, m := len(arr), len(arr[0])
	if x < 0 || x >= n || y < 0 || y >= m {
		return
	}

	// Mark the border if it's 1
	if arr[x][y] == 1 {
		arr[x][y] = 2
		visitAndMark(arr, x, y+1)
		visitAndMark(arr, x, y-1)
		visitAndMark(arr, x+1, y)
		visitAndMark(arr, x-1, y)
	}
}
