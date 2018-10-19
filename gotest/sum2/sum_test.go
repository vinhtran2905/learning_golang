package testing101

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(numbers)
		if actual != expected {
			t.Errorf(fmt.Sprintf("expected sum of %v to be %d but got %d", numbers, expected, actual))
		}
	}
}
