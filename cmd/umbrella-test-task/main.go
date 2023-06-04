package main

import (
	"log"

	"github.com/m1al04949/echo-go-server/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
