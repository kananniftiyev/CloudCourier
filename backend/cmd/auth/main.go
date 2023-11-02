package main

import (
	"backend/internal/auth/rest"
	"backend/utils"
	"log"
)

func main() {
	utils.DisplayText()
	log.Println("Auth service is starting...")
	rest.AuthStart()
}
