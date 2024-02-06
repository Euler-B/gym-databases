package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Capture connection properties
	connStr := os.Getenv("POSTGRES_URL")
	// Get a database handle
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist(`Charly Garcia`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// hard-Code ID 3 here to test the query
	album, err := albumByID(3)
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

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albums By Artist queries for albums that have the spicified artist  name.

func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows
	var albums []Album
	rows, err := db.Query("SELECT id, artist, title, price FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q:%v", name, err)
	}
	defer rows.Close()

	// Loop through rows,using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An Album to hold data from the returned row
	var alb Album

	row := db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
// func addAlbum(alb Album) (int64, error) {
// 	var //id int
// 	//result := db.QueryRow(("INSERT INTO album (title, artist, price) VALUES ($1,$2,$3) RETURNING id"), ID, Title, Price)
// 	// if err != nil {
// 	// 	return 0, fmt.Errorf("addAlbum: %v", err)
// 	// }
// 	// id, err := result.Scan()
// 	// if err != nil {
// 	// 	return 0, fmt.Errorf("addAlbum: %v", err)
// 	// }
// 	//id := result.Scan()
// 	//return id, nil
// } TODO 2: migrar la insercion de datos a consultas sql, y desacoplar el crud del main, para el proximo hito.

// TODO :
// 1.- hacer mas modular el codigo de la app
// 2.- hacer un refactor al makefile
// 3.- añadir mas consultas sql, y migraciones
// 4.- implementar tests y cobertura de los mismos
// 5.- migrar query de insert con postgres
