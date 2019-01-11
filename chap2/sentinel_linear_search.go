package main

import (
	"fmt"
)

func sentinelBetterLinearSearch(searchArray []string, arraySize int, targetName string) int {
	var last = searchArray[arraySize-1]
	searchArray[arraySize-1] = targetName
	var index = 0
	for {
		if targetName != searchArray[index] {
			index++
		} else {
			break
		}
	}

	searchArray[arraySize-1] = last
	if index < (arraySize-1) || searchArray[arraySize-1] == targetName {
		return index
	} else {
		return -1
	}

}

func main() {
	var zombies = []string{"源 さくら", "二階堂 サキ", "水野 愛", "紺野 純子", "ゆうぎり", "星川 リリィ", "山田 たえ"}

	var ret = sentinelBetterLinearSearch(zombies, len(zombies), "水野 愛")
	fmt.Println(ret)

	ret = sentinelBetterLinearSearch(zombies, len(zombies), "山田 たえ")
	fmt.Println(ret)

	ret = sentinelBetterLinearSearch(zombies, len(zombies), "種田 梨沙")
	fmt.Println(ret)
}
