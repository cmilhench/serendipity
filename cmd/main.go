package main

import (
	"log"
)

var revision = "latest"

func main() {
	log.SetFlags(log.LstdFlags)

	log.Printf("Version: (%s)\n", revision)

}
