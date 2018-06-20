package leetcode

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	input := [][2]int{
		[2]int{7, 0},
		[2]int{4, 4},
		[2]int{7, 1},
		[2]int{5, 0},
		[2]int{6, 1},
		[2]int{5, 2},
	}
	output := [][2]int{
		[2]int{5, 0},
		[2]int{7, 0},
		[2]int{5, 2},
		[2]int{6, 1},
		[2]int{4, 4},
		[2]int{7, 1},
	}
	result := reconstructQueue(input)
	if !reflect.DeepEqual(result, output) {
		t.Errorf("left: %v right: %v\n", result, output)
	}
}
