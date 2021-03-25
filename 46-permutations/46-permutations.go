package main

import (
	"reflect"
)

// https://leetcode.com/problems/permutations/

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

// Example 1:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// https://en.wikipedia.org/wiki/Backtracking
// https://medium.com/analytics-vidhya/the-blueprint-to-solve-any-backtracking-problem-b3640a3dcbd7

var permutations [][]int

func main() {
	//nums := []int{
	//	1, 2, 3,
	//}
	//
	//expected := [][]int{
	//	{1, 2, 3},
	//	{1, 3, 2},
	//	{2, 1, 3},
	//	{2, 3, 1},
	//	{3, 1, 2},
	//	{3, 2, 1},
	//}

	nums := []int{
		0, 1,
	}

	expected := [][]int{
		{0, 1},
		{1, 0},
	}

	result := permute(nums)

	if !reflect.DeepEqual(expected, result) {
		panic("failed")
	}
}

func permute(input []int) [][]int {
	var root []int
	permutations = [][]int{}

	backtracking(root, input)
	return permutations
}

func backtracking(current []int, toProcess []int) (accepted bool, result []int) {

	if accept(toProcess) {
		return true, current
	}

	for i, num := range toProcess {
		newToProcess := make([]int, len(toProcess)-1)
		copy(newToProcess, toProcess[:i])
		copy(newToProcess[i:], toProcess[i+1:])
		newCurrent := append(current, num)
		accepted, result = backtracking(newCurrent, newToProcess)
		if accepted {
			permutations = append(permutations, result)
		}
	}

	return false, current

}

func accept(toProcess []int) bool {
	return len(toProcess) == 0
}
