package main

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest"

func main() {
	rest.NewRestAdapter().Run()
}
