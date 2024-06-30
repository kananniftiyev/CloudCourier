package main

import (
	"backend/auth-service/internal/rest"
	"log"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewFigure("CloudCourier", "slant", true).Print()
	log.Println("Auth service is starting...")
	rest.AuthStart()
}
