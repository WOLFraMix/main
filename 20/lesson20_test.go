package main

import "testing"

func TestColorConstants(t *testing.T) {
	// Test the first set of color constants
	if Red != 1 {
		t.Errorf("Red = %d, want 1", Red)
	}
	if Green != 2 {
		t.Errorf("Green = %d, want 2", Green)
	}
	if Blue != 3 {
		t.Errorf("Blue = %d, want 3", Blue)
	}

	// Test the second set of color constants
	if Yellow != 0 {
		t.Errorf("Yellow = %d, want 0", Yellow)
	}
	if Orange != 1 {
		t.Errorf("Orange = %d, want 1", Orange)
	}
	if Violet != 2 {
		t.Errorf("Violet = %d, want 2", Violet)
	}
}
