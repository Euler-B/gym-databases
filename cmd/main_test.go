package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"testing"
)

var dbTest *sql.DB

type testModelAlbum struct {
	ID       int64
	Title    string
	Artist   string
	Price    float32
	Currency string
}

// add first test
func TestAlbumsByArtist(t *testing.T) {
	// Insert test data into the database
	connTestDb() 
	_, err := addTestAlbum(testModelAlbum{
		Title:    "Album 1",
		Artist:   "Artist 1",
		Price:    360.50,
		Currency: "VES",
	})

	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Call the function being tested
	albums, err := testAlbumsByArtist("Artist 1")
	if err != nil {
		t.Fatalf("Error retrieving albums by artist: %v", err)
	}

	// Check the result
	if len(albums) != 1 {
		t.Errorf("Expected 1 album, got %d", len(albums))
	}
	if albums[0].Title != "Album 1" {
		t.Errorf("Expected album title 'Album 1', got '%s'", albums[0].Title)
	}
	if albums[0].Artist != "Artist 1" {
		t.Errorf("Expected artist name 'Artist 1', got '%s'", albums[0].Artist)
	}
	if albums[0].Price != 360.50 {
		t.Errorf("Expected price 9.99, got %f", albums[0].Price)
	}
	if albums[0].Currency != "VES" {
		t.Errorf("Expected currency 'VES', got '%s'", albums[0].Currency)
	}
}


func connTestDb() {
	conStrDbTest := os.Getenv("POSTGRES_URL_TEST")
	var err error

	dbTest, err = sql.Open("postgres", conStrDbTest)
	if err != nil {
		log.Fatal(err)
	}

	pingErrTest := dbTest.Ping()
	if pingErrTest != nil {
		log.Fatal(pingErrTest)
	}

	fmt.Println("\t\t******** 🤘 DB for Testing Conected 🤘 ******* ")
}

func addTestAlbum(albtest testModelAlbum) (int64, error) {
	var idTest int64 
	err := dbTest.QueryRow("INSERT INTO album (title, artist, price, currency) VALUES ($1, $2, $3, $4) RETURNING id",
		albtest.Title, albtest.Artist, albtest.Price, albtest.Currency).Scan(&idTest)
	if err != nil {
		return 0, fmt.Errorf("addAlbumtest: %v", err)
	}
	return idTest, err

}

func testAlbumsByArtist(name string) ([]testModelAlbum, error){
	var testAlbums []testModelAlbum
	rows, err := dbTest.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albums by Artisst %q:%v", name, err)
	}
	defer rows.Close()

	for rows.Next(){
		var albtest testModelAlbum
		if err := rows.Scan(&albtest.ID, &albtest.Title, &albtest.Artist, &albtest.Price, &albtest.Currency); err != nil {
			return nil, fmt.Errorf("albums by Artist %q:%v", name, err)
		}
		testAlbums = append(testAlbums, albtest)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(" albums by artist %q:%v", name, err)
	}
	defer rows.Close()
	return testAlbums, nil
}



// TODO:
// 1.- refactorizar mis test
// 2.- hacer test unitarios
// 3.- hacer test de integracion
// 4.- crear un entorno de contenedores para tests
