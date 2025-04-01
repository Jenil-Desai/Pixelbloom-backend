package models

import "time"

type Wallpapers struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	ImageUrl     string    `json:"image_url"`
	Likes        uint64    `json:"likes"`
	Platform     string    `json:"platform"`
	ArtistsId    string    `json:"artists_id"`
	CategoriesId string    `json:"categories_id"`
	CategoryName string    `json:"category_name"`
	ArtistName   string    `json:"artist_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
