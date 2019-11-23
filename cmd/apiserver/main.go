
package main

import (
	apiserver "github.com/Andronovdima/golang-chat/internal"
	"log"
)


func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
