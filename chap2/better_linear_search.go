package main

import (
	"fmt"
)

func betterLinearSearch(searchArray []string, arraySize int, targetName string) int {
	for index := 0; index < arraySize; index++ {
		if targetName == searchArray[index] {
			return index
		}
	}
	return -1
}

func main() {
	zombies := []string{"源 さくら", "二階堂 サキ", "水野 愛", "紺野 純子", "ゆうぎり", "星川 リリィ", "山田 たえ"}

	ret := betterLinearSearch(zombies, len(zombies), "水野 愛")
	fmt.Println(ret)

	ret = betterLinearSearch(zombies, len(zombies), "種田 梨沙")
	fmt.Println(ret)
}
