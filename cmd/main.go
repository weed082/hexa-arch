package main

import (
	"fmt"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
)

func main() {

	repository := repository.New()

	defer repository.Disconnect()

	result, err := repository.Test()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
