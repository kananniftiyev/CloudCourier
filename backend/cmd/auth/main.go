package main

import (
	"backend/internal/auth/rest"
	"github.com/common-nighthawk/go-figure"
	"log"
)

func DisplayText() {
	figure.NewFigure("CloudShareX", "slant", true).Print()
}

func main() {
	DisplayText()
	log.Println("Auth service is starting...")
	rest.AuthStart()
}
