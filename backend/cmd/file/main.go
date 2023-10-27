package main

import (
	"backend/internal/file/rest"
	"github.com/common-nighthawk/go-figure"
	"log"
)

func DisplayText() {
	figure.NewFigure("CloudShareX", "slant", true).Print()
}

func main() {

	DisplayText()
	log.Println("File service is starting...")
	rest.FileStart()

}
