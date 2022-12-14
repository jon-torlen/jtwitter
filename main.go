package main

import (
	"log"

	"github.com/jon-torlen/jtwitter/bd"
	"github.com/jon-torlen/jtwitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conección a la BD")
		return
	}
	handlers.Manejadores()

}
