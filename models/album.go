package models

type Album struct {
	ID       int64
	Title    string
	Artist   string
	Price    float32
	Currency string
}

// TODO: parse struct's field with json.
