package task_1

import "strings"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 将第一个字符串作为前缀，逐个比较，如果有不是共同前缀的则缩短前缀
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pre) {
			if len(pre) == 0 {
				return ""
			}
			pre = pre[:len(pre)-1]
		}
	}

	return pre
}
