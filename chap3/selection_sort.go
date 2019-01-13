package main

import "fmt"

func selectionSort(array []string, arraySize int) []string {
	for index := 0; index < arraySize-1; index++ {
		smallest := array[index]
		smallestIndex := index
		fmt.Println("first smallest", smallest)
		for j := index; j < arraySize; j++ {
			if smallest > array[j] {
				fmt.Println("change smallest", smallest, "->", array[j])
				smallest = array[j]
				smallestIndex = j
			}
		}
		fmt.Println("end smallest", smallest)
		array[smallestIndex] = array[index]
		array[index] = smallest

	}
	return array
}

func main() {
	unsortArray := []string{"f", "g", "e", "d", "c", "b", "a"}
	fmt.Println(selectionSort(unsortArray, len(unsortArray)))
}
