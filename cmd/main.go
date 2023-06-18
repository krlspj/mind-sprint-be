package main

import (
	"log"

	"github.com/krlspj/mind-sprint-be/cmd/bootstrap"
)

const applicationName = "Mind Sprint"

func main() {
	log.Printf("Starting %s backend", applicationName)

	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
