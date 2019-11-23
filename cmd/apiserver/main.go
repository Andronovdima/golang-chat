package main

import (
	apiserver "Golang/FworkChat/golang-chat/internal"
	"log"
)


func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}

