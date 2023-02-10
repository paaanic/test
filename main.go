package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findExprs(nums []int, target int) []string {
	res := make([]string, 0)

	var backtrack func(int, int, int, []string)
	backtrack = func(i, op, v int, expr []string) {
		if i == len(nums) {
			if v == target && op == 0 {
				res = append(res, strings.Join(expr[1:], ""))
			}
			return
		}

		for i := nDigits(nums[i]); i > 0; i-- {
			op *= 10
		}
		op += nums[i]
		opS := strconv.FormatInt(int64(op), 10)

		if op != 0 {
			backtrack(i+1, op, v, expr)
		}

		backtrack(i+1, 0, v+op, append(expr, "+", opS))
		if i > 0 {
			backtrack(i+1, 0, v-op, append(expr, "-", opS))
		}
	}

	backtrack(0, 0, 0, []string{})

	return res
}

func nDigits(a int) int {
	if a == 0 {
		return 1
	}

	c := 0
	for a != 0 {
		a /= 10
		c++
	}

	return c
}

func main() {
	for _, expr := range findExprs([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 200) {
		fmt.Println(expr)
	}

	/*
		Output:

		98+76-5+43-2-10
		98-7+65+43+2-1+0
		98-7+65+43+2-1-0
		98+76+5+4+3+210
		9-8+7-6-5-4-3+210
		9-8-7-6-5+4+3+210
	*/
}
