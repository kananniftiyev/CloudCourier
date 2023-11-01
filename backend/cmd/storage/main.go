package main

import (
	"backend/internal/storage"
	"backend/utils"
	"log"
)

func main() {
	utils.DisplayText()
	log.Println("Storage service is starting...")
	storage.StartStorageCheck()
}
