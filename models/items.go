package models

// Item ... One would wish to perform in their lifetime.
type Item struct {

	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type allItems []Item

// Items ... collections of all items
var Items = allItems{
		Item{1, "Visit the notre dam", "would be ideal to visit the notre dam this year"},
		Item{2, "Visit Dian", " visit the dian beach at the end of this year "},
	}