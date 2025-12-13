package main

import (
	"MateNode-Golang/task_1"
	"fmt"
)

func main() {
	num := 313
	barcket_s := "{}()[]"
	barcket_sj_false := "{}()[](}"

	fmt.Println(task_1.IsPalindrome(num))
	fmt.Println(task_1.IsValid(barcket_s))
	fmt.Println(task_1.IsValid(barcket_sj_false))

}
