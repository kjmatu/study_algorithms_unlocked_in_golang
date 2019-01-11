package main

import "fmt"

func procedureFactorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * procedureFactorial(n-1)
	}
}

func main() {
	fmt.Println(procedureFactorial(5))
}
