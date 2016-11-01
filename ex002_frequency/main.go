package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 10, 3, 5, 7, 1, 2, 4, 10, 9, 9, 1, 4, 6, 1, 7, 10, 9, 6, 5, 1, 4, 10, 1, 7, 7, 5, 6, 7, 6, 7, 9, 2, 6, 4, 5, 1, 8, 8, 3, 3, 7, 2, 4, 1, 5, 9, 1, 1, 5, 9, 3, 6, 7, 1, 4, 1, 3, 2, 2, 7, 4, 7, 7, 7, 7, 10, 2, 2, 10, 6, 6, 5, 9, 1, 10, 3, 4, 8, 8, 7, 9, 8, 3, 10, 4, 8, 9, 1, 9, 6, 9, 2, 6, 1, 8, 2, 7}

	frequency, uniques := analyseFrequency(nums)

	fmt.Println(frequency) // expected: map[1:14 4:9 9:11 5:8 2:10 3:7 7:15 6:10 8:7 10:9]
	fmt.Println(uniques)   // expected: [10 5 2 3 7 1 4 9 6 8]  --- in the order in which they occur in nums
}

func analyseFrequency(nums []int) (map[int]int, []int) {

	uniques := getUniques(nums)
	frequency := toSortedMap(bubbleSort(nums))



	return frequency, uniques
}

func getUniques(input []int) []int {
	uniques := make([]int,0)
	for _, val := range input {
		if (!valueInSlice(val, uniques)) {
			uniques = append(uniques, val)
		}
	}
	return uniques
}

func valueInSlice(val int, list []int) bool {
	for _, cur := range list {
		if cur == val {
			return true
		}
	}
	return false
}

func toSortedMap(input []int) (map[int]int) {
	frequency := make(map[int]int)

	lastInput := 0
	for i:=0; i<len(input); i++ {
		if(input[i] != lastInput) {
			frequency[input[i]] = 1
		} else {
			frequency[input[i]] = frequency[input[i]] + 1
		}
		lastInput = input[i]
	}

	return frequency
}

func bubbleSort(input []int) []int {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n-1; i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

