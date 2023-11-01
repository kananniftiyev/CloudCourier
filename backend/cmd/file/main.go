package main

import (
	"backend/internal/file/rest"
	"backend/utils"
	"log"
)

func main() {

	utils.DisplayText()
	log.Println("File service is starting...")
	rest.FileStart()

}
