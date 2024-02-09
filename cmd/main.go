package main

import (
	"fmt"
	"log"

	"github.com/euler-b/access-relational-database/internal/database"
	
)

func main() {
	database.ConnectToDb()
	
	// albums, err := database.AlbumsByArtist("Charly Garcia")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// hard-Code ID 3 here to test the query
	album, err := database.AlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album found: %v\n", album)

	// // Hard-code data into db
	// albID, err := addAlbum(Album{
	// 	Title:  "Big Bang",
	// 	Artist: "Enanitos Verdes",
	// 	Price:  6.50,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of album added: %v\n", albID)
}

// TODO :
// 1.- hacer mas modular el codigo de la app
// 2.- hacer un refactor al makefile
// 3.- añadir mas consultas sql, y migraciones
// 4.- implementar tests y cobertura de los mismos
// 5.- migrar query de insert con postgres
