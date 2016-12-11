package main

import (
	"github.com/likitsaelee/gotodo/router"
	"log"
)

func main() {
	r := router.Load()
	log.Println("Running on port 8080...")
	r.Run()
}
