package task_1

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/
func PlusOne(digits []int) []int {
	n := len(digits)
	// 从最低位开始加 1
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 { // 无进位，提前结束
			return digits
		}
		digits[i] = 0 // 有进位，本位变 0 继续循环
	}
	// 走到这里说明全是 999...，需要扩容
	ans := make([]int, n+1)
	ans[0] = 1 // 其余位默认 0
	return ans
}
