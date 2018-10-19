package testing101

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

func ExampleSum() {
	numbers := []int{4, 4, 4}
	fmt.Println(Sum(numbers))
	//Output:
	//14
}

func TestSumWithNegatice(t *testing.T) {

}
