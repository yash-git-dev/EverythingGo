package Controllers

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// Initialize left and right arrays
	left := make([]int, n)
	right := make([]int, n)

	// Initialize result array
	answer := make([]int, n)

	// Calculate left products
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// Calculate right products
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// Calculate final result
	for i := 0; i < n; i++ {
		answer[i] = left[i] * right[i]
	}

	return answer
}

func SelfMultiplierMain() {
	// Example 1
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // Output: [24 12 8 6]

	// Example 2
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2)) // Output: [0 0 9 0 0]
}
