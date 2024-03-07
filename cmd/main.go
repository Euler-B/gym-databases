package main

import (
	"fmt"
	"log"

	"github.com/euler-b/access-relational-database/internal/database"
	"github.com/euler-b/access-relational-database/models"
)

func main() {
	database.ConnectToDb()
	
	albums, err := database.AlbumsByArtist("Charly Garcia")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	//hard-Code ID 2 here to test the query
	album, err := database.AlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album found: %v\n", album)

	//Hard-code data into db
	albID, err := database.AddAlbum(models.Album{
		Title:  "Vasos y besos",
		Artist: "Los Abuelos de la nada",
		Price:  5943.97,
		Currency: "ARS",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of album added: %v\n", albID)
}

// TODO :
// 1.- hacer mas modular el codigo de la app
// 2.- hacer un refactor al makefile
// 3.- añadir mas consultas sql, y migraciones
// 4.- implementar tests y cobertura de los mismos
// 5.- migrar query de insert con postgres
