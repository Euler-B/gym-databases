package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/euler-b/access-relational-database/models"

	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDb() {
	connStr := os.Getenv("POSTGRES_URL")
	// Get a database handle
	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	//msg = "******** DB Conected *****"
	fmt.Println("\t\t******** 🤘 DB Conected 🤘 ******* ")

}

// func init() {
// 	err := godotenv.Load("internal/database/.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// albums By Artist queries for albums that have the spicified artist  name.

func AlbumsByArtist(name string) ([]models.Album, error) {

	// An albums slice to hold data from returned rows
	var albums []models.Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q:%v", name, err)
	}
	defer rows.Close()

	// Loop through rows,using Scan to assign column data to struct fields
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.Currency); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func AlbumByID(id int64) (models.Album, error) {
	// An Album to hold data from the returned row
	var alb models.Album

	row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.Currency); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsByID %d: %v", id, err)
	}

	return alb, nil
}

// addAlbum adds the specified album to the	 database,
// returning the album ID of the new entry
func AddAlbum(alb models.Album) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO album (title, artist, price, currency) VALUES ($1, $2, $3, $4) RETURNING id",
		alb.Title, alb.Artist, alb.Price, alb.Currency).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, err
}

// TODO 2: migrar la insercion de datos a consultas sql, y desacoplar el crud del main, para el proximo hito.
