package main

import (
	"github.com/Parshva28/github-action/microservice"
	"log"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
