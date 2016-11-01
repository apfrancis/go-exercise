package main

import "testing"

func TestTrecimateReturnOne (t *testing.T) {
	if(runTestWithNumber(33) != 1) {
		t.Error("Expected 1, got ", runTestWithNumber(33))
	}
}

func runTestWithNumber (num int) int {
	return trecimate(num)
}