package main

import (
	"backend/file-service/internal/rest"
	"log"

	"github.com/common-nighthawk/go-figure"
)

func main() {

	figure.NewFigure("CloudShareX", "slant", true).Print()
	log.Println("File service is starting...")
	rest.FileStart()

}
