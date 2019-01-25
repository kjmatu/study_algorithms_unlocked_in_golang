package main

import "fmt"

func mergeGoLike(left, right []string) []string {
	var merged []string
	li, ri := 0, 0

	// ソート済み配列をマージするため、それぞれ左から見ていくだけで良い
	for li < len(left) && ri < len(right) {
		if left[li] <= right[ri] {
			merged = append(merged, left[li])
			li++
		} else {
			merged = append(merged, right[ri])
			ri++
		}
		// fmt.Println("merged", merged)
	}

	// 上のwhile文のどちらかがFalseになった場合終了するため、あまりをextendする
	if li < len(left) {
		merged = append(merged, left[li:]...)
		li++
	} else {
		merged = append(merged, right[ri:]...)
		ri++
	}
	return merged
}

func mergeSortGoLike(array []string) []string {
	if len(array) <= 1 {
		return array
	}

	// ここで分割を行う
	mid := len(array) / 2
	// fmt.Println("mid", mid)

	// 再帰的に分割を行う
	left := array[:mid]
	right := array[mid:]
	// fmt.Println("left", left)
	// fmt.Println("right", right)

	left = mergeSortGoLike(left)
	right = mergeSortGoLike(right)

	// returnが返ってきたら、結合を行い、結合したものを次に渡す
	return mergeGoLike(left, right)
}

func main() {
	unsortArray := []string{"h", "g", "f", "e", "d", "c", "b", "a"}
	fmt.Println(unsortArray)
	fmt.Println(len(unsortArray))
	sortArray := mergeSortGoLike(unsortArray)
	fmt.Println(sortArray)
	fmt.Println(len(sortArray))
}
