package main

import (
	"learngo/algo/utils"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	left, right := 0, 0
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= s {
			ans = utils.Min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
