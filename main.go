package main

import (
	"fmt"
	"log"
	"net/http"
	"pokemon-list/injector"
)

func CheckPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}

func main() {
	CheckPrime(0)  // not a prime
	CheckPrime(1)  // not a prime
	CheckPrime(3)  // a prime
	CheckPrime(71) // not a prime

	s := injector.InitializedServer()
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
