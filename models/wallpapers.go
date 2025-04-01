package models

import "time"

type Wallpapers struct {
	Id           string
	Title        string
	ImageUrl     string
	Likes        uint64
	Platform     string
	CategoriesId string
	ArtistId     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
