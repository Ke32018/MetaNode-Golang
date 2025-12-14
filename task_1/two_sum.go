package task_1

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums)) // 值 -> 下标
	for i, v := range nums {
		if j, ok := seen[target-v]; ok { // 找到配对
			return []int{j, i}
		}
		seen[v] = i // 记录自己的下标
	}
	return nil
}
