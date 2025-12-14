package task_1

import (
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 1. 按左端点升序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	curStart, curEnd := intervals[0][0], intervals[0][1]

	// 2. 扫描合并
	for i := 1; i < len(intervals); i++ {
		s, e := intervals[i][0], intervals[i][1]
		if s <= curEnd { // 重叠，合并
			curEnd = max(curEnd, e) // 右端点取最大
		} else { // 无法合并，存档
			merged = append(merged, []int{curStart, curEnd})
			curStart, curEnd = s, e // 开启新区间
		}
	}

	merged = append(merged, []int{curStart, curEnd})
	return merged
}
