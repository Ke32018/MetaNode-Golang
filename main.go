package main

import (
	"MateNode-Golang/task_1"
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	// digits := []int{1, 2, 3}
	digits2 := []int{9}

	fmt.Println(task_1.LongestCommonPrefix(strs))
	// fmt.Println(task_1.PlusOne(digits))
	fmt.Println(task_1.PlusOne(digits2))
}
