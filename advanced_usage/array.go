package advanced_usage

import "fmt"

func Array() {
	a := [4]int{4, 3, 2, 1}

	// 方式1，使用下标读取数据
	element := a[2]
	fmt.Println("element = ", element)

	// 方式2，使用range遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
	}

	for i := range a {
		fmt.Println("only index, index = ", i)
	}

	// 读取数组长度
	fmt.Println("len(a) = ", len(a))
	// 使用下标，for循环遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println("use len(), index = ", i, "value = ", a[i])
	}

	// 数组的部分特性类似基础数据类型，当数组作为参数传递时，在函数中并不能改变外部实参的值。

	modifyArray(a)
	fmt.Println("after modify, a = ", a)

	// 如果想要修改外部实参的值，需要把数组的指针作为参数传递给函数。
	modifyArrayPointer(&a)
	fmt.Println("after modify pointer, a = ", a)

	printFuncParamPointer(carr)
}

func modifyArray(a [4]int) {
	a[0] = 100
	fmt.Println("in modifyArray func, a = ", a)
}

func modifyArrayPointer(a *[4]int) {
	a[0] = 100
}

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}
}
