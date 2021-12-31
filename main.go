package main

import (
	"log"
	"net/http"
	"pokemon-list/injector"
)

func main() {
	s := injector.InitializedServer()
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
