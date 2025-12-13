package advanced_usage

import "fmt"

func MultiArray() {
	// 二维数组
	a := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Println("a = ", a)

	// 三维数组
	b := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Println("b = ", b)

	// 也可以省略各个位置的初始化,在后续代码中赋值
	c := [3][3][3]int{}
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Println("c = ", c)
}
