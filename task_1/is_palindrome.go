package task_1

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 数学解法：将数字后半部分反转，然后与前半部分比较
	rev := 0
	for x > rev { // 当 x < rev 时，处理已过半
		rev = rev*10 + x%10
		x /= 10
	}

	return x == rev || x == rev/10 // 如果数字长度为奇数，则需要去掉中间的数字
}
