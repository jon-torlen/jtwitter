package main

import (
	"log"

	"github.com/jon-torlen/jtwitter.git/bd"
	"github.com/jon-torlen/jtwitter.git/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conección a la BD")
		return
	}
	handlers.Manejadores()

}
