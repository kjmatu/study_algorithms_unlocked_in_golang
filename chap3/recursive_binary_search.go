package main

import "fmt"

func recursiveBinarySearch(searchArray []int, startIndex int, endIndex int, targetNumber int) int {
	if startIndex > endIndex {
		return -1
	} else {
		centerIndex := (startIndex + endIndex) / 2
		if searchArray[centerIndex] == targetNumber {
			return centerIndex
		} else {
			if searchArray[centerIndex] > targetNumber {
				return recursiveBinarySearch(searchArray, startIndex, centerIndex-1, targetNumber)
			} else {
				return recursiveBinarySearch(searchArray, centerIndex+1, endIndex, targetNumber)
			}
		}
	}
}

func main() {
	array := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 16}
	ret := recursiveBinarySearch(array, 0, len(array)-1, 17)
	fmt.Println(ret)
}
