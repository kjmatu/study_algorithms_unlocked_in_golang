package main

import "fmt"

func recursiveLinearSearch(searchArray []string, arraySize int, index int, targetName string) int {
	// fmt.Println(index)
	// fmt.Println(arraySize)
	if index > (arraySize - 1) {
		return -1
	} else if (index <= arraySize) && (searchArray[index] == targetName) {
		return index
	} else if (index <= arraySize) && (searchArray[index] != targetName) {
		return recursiveLinearSearch(searchArray, arraySize, index+1, targetName)
	}
	return -1
}

func main() {
	var zombies = []string{"源 さくら", "二階堂 サキ", "水野 愛", "紺野 純子", "ゆうぎり", "星川 リリィ", "山田 たえ"}

	var ret = recursiveLinearSearch(zombies, len(zombies), 0, "水野 愛")
	fmt.Println(ret)

	ret = recursiveLinearSearch(zombies, len(zombies), 0, "山田 たえ")
	fmt.Println(ret)

	ret = recursiveLinearSearch(zombies, len(zombies), 0, "種田 梨沙")
	fmt.Println(ret)

}
