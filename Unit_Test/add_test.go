package main

import "testing"

func TestAddCases(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 0},
	}

	for _, c := range cases {
		result := add(c.a, c.b)
		if result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}

	}
}
