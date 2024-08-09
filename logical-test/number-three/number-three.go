package main

import "fmt"

func generateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	}

	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		next := fib[i-1] + fib[i-2]
		fib = append(fib, next)
	}

	return fib
}

func main() {
	n := 10
	fib := generateFibonacci(n)

	for _, value := range fib {
		fmt.Printf("%d ", value)
	}
}
