package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)

	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Error("Result is not correct")
	}
}
