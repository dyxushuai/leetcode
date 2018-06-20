package leetcode

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	input := "ababcbacadefegdehijhklij"
	output := []int{9, 7, 8}
	result := partitionLabels(input)
	if !reflect.DeepEqual(result, output) {
		t.Errorf("left: %v right: %v\n", result, output)
	}
}
