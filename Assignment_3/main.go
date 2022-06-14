package main

import (
	"Assignment_3/app"
	"fmt"
)

func main() {

	fmt.Println(app.CheckDuplicates([]int{1, 1, 1, 2, 3, 3, 1}))
	fmt.Println(app.CheckDuplicatesBruteForce([]int{1, 1, 1, 2, 3, 3, 1}))
	fmt.Println(app.TargetSum([]int{3, 4, 2, 3, 6, 5}, 9))
}
