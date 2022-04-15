package main

import (
	"https://github.com/Parshva28/github-action/tree/main/microservice"
	"log"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
