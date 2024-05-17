package main

import (
	"fmt"
	"sort"
)

func result() string {
	var n int
	fmt.Scan(&n)

	containers := make([][]int, n)
	for i := range containers {
		containers[i] = make([]int, n)
	}
	col := make([]int, n)
	row := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&containers[i][j])
			col[j] += containers[i][j]
			row[i] += containers[i][j]
		}
	}
	sort.Ints(row)
	sort.Ints(col)

	for i := 0; i < n; i++ {
		if row[i] != col[i] {
			return "no"
		}
	}

	return "yes"
}
func main() {

	res := result()
	fmt.Println(res)

}
