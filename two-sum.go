func twoSum(nums []int, target int) []int {
	for i, _ := range nums {

		for j := 1 + i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return []int{}
}
