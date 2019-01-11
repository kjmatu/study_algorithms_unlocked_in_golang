package main

import "fmt"

func procedureBadFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return procedureBadFactorial((n + 1) / (n + 1))
	}
}

func main() {
	fmt.Println(procedureBadFactorial(5))
}
