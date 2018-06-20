package leetcode

import (
	"fmt"
	"sort"
)

// A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

// Example 1:
// Input: S = "ababcbacadefegdehijhklij"
// Output: [9,7,8]
// Explanation:
// The partition is "ababcbaca", "defegde", "hijhklij".
// This is a partition so that each letter appears in at most one part.
// A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
// Note:

// S will have length in range [1, 500].
// S will consist of lowercase letters ('a' to 'z') only.

func partitionLabels(input string) []int {
	// [2]int{start, end}
	charAppear := map[int32][2]int{}
	for idx, char := range input {
		if v, ok := charAppear[char]; ok {
			if v[1] < idx {
				charAppear[char] = [2]int{v[0], idx}
			}
		} else {
			charAppear[char] = [2]int{idx, idx}
		}
	}
	appearSlice := [][2]int{}
	for _, v := range charAppear {
		appearSlice = append(appearSlice, v)
	}
	// sort by start position
	sort.Slice(appearSlice, func(i, j int) bool {
		return appearSlice[i][0] < appearSlice[j][0]
	})
	fmt.Println(appearSlice)
	result := []int{}
	sum := 0
	end := appearSlice[0][1]
	for _, v := range appearSlice {
		if v[0] > end {
			if len(result) == 0 {
				cur := end + 1
				result = append(result, cur)
				sum += cur
			} else {
				prev := result[len(result)-1]
				cur := end + 1 - prev
				result = append(result, cur)
				sum += cur
			}
		}
		if v[1] > end {
			end = v[1]
		}
	}
	if sum < len(input) {
		result = append(result, len(input)-sum)
	}
	return result
}
