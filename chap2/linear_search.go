package main

import "fmt"

func linear_search(search_array []string, array_size int, target_name string) int8 {
	var answer_index int8
	answer_index = -1
	for index := 0; index < array_size; index++ {
		if search_array[index] == target_name {
			answer_index = int8(index)
			return answer_index
		}
	}
	return answer_index
}

func main() {
	name_array := []string{"源 さくら", "二階堂 サキ", "水野 愛", "紺野 純子", "ゆうぎり",
		"星川 リリィ", "山田 たえ"}

	ret := linear_search(name_array, len(name_array), "水野 愛")
	fmt.Println(ret)

	ret = linear_search(name_array, len(name_array), "田村 ゆかり")
	fmt.Println(ret)
}
