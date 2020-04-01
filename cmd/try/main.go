package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("KLOOK_RECOMMENDATION_BIGQUERY_CREDENTIAL"))
}

func Fib(n int) int {
	if n <= 2 {
		return n
	}

	return Fib(n-1) * Fib(n-2)
}
