package main

import "fmt"

func binarySearch(searchArray []int, arraySize int, targetNumber int) int {
	startIndex := 1
	endIndex := arraySize - 1

	for startIndex <= endIndex {
		centerIndex := (startIndex + endIndex) / 2
		if searchArray[centerIndex] == targetNumber {
			return centerIndex
		} else {
			if searchArray[centerIndex] > targetNumber {
				endIndex = centerIndex - 1
			} else {
				startIndex = centerIndex + 1
			}
		}
	}
	return -1
}

func main() {
	array := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 16}
	ret := binarySearch(array, len(array), 16)
	fmt.Println(ret)
}
