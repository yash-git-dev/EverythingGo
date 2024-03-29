// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]
package Controllers

import (
	"errors"
	"sort"
)

func FrequentInteger(nums []int, k int) ([]int, error) {
	//final := make([]int, k)
	count := make(map[int]int)

	if len(nums) == 0 {
		return []int{}, errors.New("empty array provided")
	}

	for _, no := range nums {

		count[no]++

	}
	var uniqueNums []int
	for num := range count {
		uniqueNums = append(uniqueNums, num)
	}

	// Sort the unique numbers based on their frequency in descending order
	sort.Slice(uniqueNums, func(i, j int) bool {
		return count[uniqueNums[i]] > count[uniqueNums[j]]
	})

	// Take the first k elements from the sorted slice
	var result []int
	for i := 0; i < k && i < len(uniqueNums); i++ {
		result = append(result, uniqueNums[i])
	}

	return result, nil
}
