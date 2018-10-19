package sum_test

import (
	"gopractice/gotest/sum"
	"testing"
)

func TestInts(t *testing.T) {

	// t.Errorf("This test failed becase I said so")
	// t.Fatalf("This test failed and stopped running")

	s := sum.Ints(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("failed got %v, expected 15", s)
	}

	// s = sum.Ints()
	// if s != 0 {
	// 	t.Errorf("failed, got %v, expected 0", s)
	// }
	//
	// s = sum.Ints(1, -1)
	// if s != 0 {
	// 	t.Errorf("failed, got %v, expected 0", s)
	// }

	tt := []struct {
		title_tc    string
		numbers     []int
		expectedsum int
	}{
		{"TC1: one to five", []int{1, 2, 3, 4, 5}, 15},
		// {"TC2:no numbers", nil, 0},
		// {"TC3: one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		total := sum.Ints(tc.numbers...)
		if total != tc.expectedsum {
			t.Errorf("%v Failed, got %v, expected %v", tc.title_tc, total, tc.expectedsum)
		}
	}
}
