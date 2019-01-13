package main

import "fmt"

func insertion_sort(array []string, arraySize int) []string {
	// 2番目の配列要素から挿入ソートを行う
	index := 1

	for index < arraySize {
		// 未ソートの要素をkeyにする
		key := array[index]
		// ソート済み配列の一番右のindexをjに格納
		j := index - 1

		for {
			// 未ソート要素がソート済み配列の一番右より小さかったら
			if j >= 0 && array[j] > key {
				// ソート済み配列の要素を1個右へずらす
				array[j+1] = array[j]
				// ソート済み配列の一番右のindexを1個左に移動
				j = j - 1
			} else {
				break
			}
		}
		// 上記の移動で空白になったindexにkeyを格納
		array[j+1] = key
		index++
	}
	return array
}

func main() {
	unsortArray := []string{"f", "g", "e", "d", "c", "b", "a"}
	fmt.Println(insertion_sort(unsortArray, len(unsortArray)))
}
