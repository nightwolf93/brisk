package main

import (
	"log"

	"github.com/nightwolf93/brisk/api"
)

// version is the string ref of the version of brisk
var version = "0.0.1"

func main() {
	log.Printf("starting brisk v%s ..", version)
	api.Init()
}
