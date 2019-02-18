package main

import "testing"

func TestGetAverage(t *testing.T) {
	var a float64
	a = getAverage(1, 2)
	if a != 1.5 {
		t.Error("Expected 1.5, got ", a)
	}
}
