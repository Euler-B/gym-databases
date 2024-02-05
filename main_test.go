package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

// add first test
func TestAlbumsByArtist(t *testing.T) {
	setupDB()
	// Insert test data into the database
	_, err := addAlbum(Album{
		Title:  "Album 1",
		Artist: "Artist 1",
		Price:  9.99,
	})
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Call the function being tested
	albums, err := albumsByArtist("Artist 1")
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
	if albums[0].Price != 9.99 {
		t.Errorf("Expected price 9.99, got %f", albums[0].Price)
	}
}

func setupDB() { 
	cfg := mysql.Config {
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   "172.20.0.1:3306",
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}
	
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
}

// TODO:
// 1.- refactorizar mis test
// 2.- hacer test unitarios
// 3.- hacer test de integracion
// 4.- crear un entorno de contenedores para tests