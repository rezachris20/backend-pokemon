package main

import (
	"fmt"
	"log"
	"net/http"
	"pokemon-list/injector"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {

	for i := 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
	}
	fmt.Println("")

	s := injector.InitializedServer()
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
