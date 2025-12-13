package advanced_usage

import "fmt"

// 可以使用 make 函数来创建 map``
// make(map[key-type]value-type, [capacity])
// capacity 是可选的，用于指定 map 的初始容量。
// 如果 capacity 大于 0，则 map 会预分配足够的内存空间，以提高性能。
// 如果 capacity 小于等于 0，则 map 会根据实际情况动态分配内存空间。

func TestMap() {
	var m1 map[string]string
	fmt.Println("m1 length:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 length:", len(m2))
	fmt.Println("m2 =", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println("m3 =", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)

	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)
}
