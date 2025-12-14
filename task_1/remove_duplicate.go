package task_1

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。

nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1 // 慢指针：下一个唯一元素该放的位置
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // 遇到新数字
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
