package main

import (
	"testing"
	"reflect"
)

func TestReturnsCorrectFrequenciesAndUniques (t *testing.T) {
	nums := []int {1, 1, 1, 1, 1, 1}
	analyseFrequency(nums)
	frequency, uniques := analyseFrequency(nums)
	if(frequency[1] != 6) {
		t.Fail()
	}
	if(len(uniques) != 1) {
		t.Fail()
	}
}

func TestBubbleSortSortsCorrectly (t *testing.T) {
	nums := []int {3, 4, 2, 1, 7, 8}
	expected := []int {1, 2, 3, 4, 7, 8}
	if(!reflect.DeepEqual(bubbleSort(nums),expected)) {
		t.Fail()
	}
}

func TestValueInSliceReturnsTrueIfValueExists  (t *testing.T) {
	slice := []int {1,2,3,4}
	if(!valueInSlice(3,slice)) {
		t.Fail()
	}
}

func TestValueInSliceReturnsFalseIfValueDoesNotExist  (t *testing.T) {
	slice := []int {1,2,3,4}
	if(valueInSlice(10,slice)) {
		t.Fail()
	}
}