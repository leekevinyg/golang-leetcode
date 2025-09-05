package problems

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		value, exists := seen[complement]
		if exists {
			ans := []int{i, value}
			return ans
		} else {
			seen[nums[i]] = i
		}
	}
	return []int{0, 0}
}
