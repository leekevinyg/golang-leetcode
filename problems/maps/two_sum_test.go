package problems

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("basic example 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("basic example 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("same number twice", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		nums := []int{-1, -2, -3, -4, -5}
		target := -8
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("mixed positive and negative", func(t *testing.T) {
		nums := []int{-3, 4, 3, 90}
		target := 0
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("zero target", func(t *testing.T) {
		nums := []int{0, 4, 3, 0}
		target := 0
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("large numbers", func(t *testing.T) {
		nums := []int{1000000, 2000000, 3000000}
		target := 5000000
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("first two elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("last two elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 9
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("no solution exists", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		target := 10
		result := TwoSum(nums, target)
		expected := []int{0, 0}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("TwoSum(%v, %d) = %v, expected %v when no solution exists",
				nums, target, result, expected)
		}
	})

	t.Run("duplicate numbers with solution", func(t *testing.T) {
		nums := []int{3, 5, 3, 4}
		target := 6
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("minimum array size", func(t *testing.T) {
		nums := []int{1, 5}
		target := 6
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})

	t.Run("edge case with zeros", func(t *testing.T) {
		nums := []int{0, 0, 3, 4}
		target := 0
		result := TwoSum(nums, target)

		if !isValidTwoSumResult(nums, target, result) {
			t.Errorf("TwoSum(%v, %d) = %v, but nums[%d] + nums[%d] != %d",
				nums, target, result, result[0], result[1], target)
		}
	})
}

// Helper function to validate that the result indices point to numbers that sum to target
func isValidTwoSumResult(nums []int, target int, result []int) bool {
	if len(result) != 2 {
		return false
	}

	i, j := result[0], result[1]

	// Check bounds
	if i < 0 || i >= len(nums) || j < 0 || j >= len(nums) {
		return false
	}

	// Check that indices are different
	if i == j {
		return false
	}

	// Check that the sum equals target
	return nums[i]+nums[j] == target
}
