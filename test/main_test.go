package main

import "testing"

func TestGetAverage(t *testing.T) {
	var a float64
	a = getAverage(1, 2)
	if a != 1.5 {
		t.Error("Expected 1.5, got ", a)
	}
}

func TestGetSum(t *testing.T) {
	var s int
	s = getSum(1,2,3,4)
	if s != 10 {
		t.Error("Expected 10, got ", s)
	}
}