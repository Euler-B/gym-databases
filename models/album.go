package models

type Album struct {
	Id     int64
	Title  string
	Artist string
	Price  float32
}

// TODO: parse struct's field with json.
