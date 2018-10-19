package main

import "testing"

// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
//
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got %d, expected %d", total, 10)
// 	}
// }

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
		{2, -1, 1},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.z {
			t.Errorf("Sum of (%d + %d) was incorrect, got %d, expected %d", table.x, table.y, total, table.z)
		}
	}
}
