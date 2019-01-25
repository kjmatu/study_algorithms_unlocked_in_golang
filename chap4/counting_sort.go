package main

import "fmt"

func countKeysEqual(array []uint, maxValue uint) []uint {
	equal := make([]uint, maxValue)
	for i := 0; i < len(array); i++ {
		key := array[i]
		equal[key]++
	}
	return equal
}

func countKeysLess(equalArray []uint, equalArrayMaxValue uint) []uint {
	less := make([]uint, equalArrayMaxValue)
	for j := 1; j < int(equalArrayMaxValue); j++ {
		less[j] = less[j-1] + equalArray[j-1]
	}
	return less
}

func rearrange(array []uint, lessArray []uint, maxValue uint) []uint {
	sortArray := make([]uint, len(array))
	fmt.Println("sortArray", sortArray)

	// next := make([]uint, maxValue)
	// for j := 0; j < len(next); j++ {
	// 	next[j] = lessArray[j]
	// }
	// fmt.Println("next", next)
	next := lessArray
	for i := 0; i < len(sortArray); i++ {
		key := array[i]
		index := next[key]
		// fmt.Println("i", i)
		// fmt.Println("key", key)
		// fmt.Println("index", index)
		// fmt.Println("array[i]", array[i])
		// fmt.Println("sortArray[index]", sortArray[index])
		sortArray[index] = array[i]
		next[key]++
	}
	return sortArray
}

func main() {
	sampleArray := []uint{4, 1, 5, 0, 1, 6, 5, 1, 5, 3}
	equalArray := countKeysEqual(sampleArray, 7)
	fmt.Println(equalArray)
	fmt.Println(len(equalArray))
	lessArray := countKeysLess(equalArray, 7)
	fmt.Println(lessArray)
	fmt.Println(len(lessArray))
	sortArray := rearrange(sampleArray, lessArray, 7)
	fmt.Println(sortArray)
}
